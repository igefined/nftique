package rate_limiter

import (
	"context"
	"crypto/md5"
	"fmt"
	"math/big"
	"os"
	"sync"
	"time"

	cfg "github.com/igefined/nftique/pkg/config"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

const (
	Auth   RequestType = "auth"
	Common RequestType = "common"

	Undefined RequestType = "undefined"
)

const defaultTokenSize = 5

type RequestType string

type RateLimiter struct {
	mu      sync.RWMutex
	buckets map[string]*TokenBucket

	logger   *zap.Logger
	settings *Settings

	redisClient *redis.Client
	redisScript *redis.Script
}

func NewRateLimiter(
	cfg *cfg.RedisCfg,
	rlCfg *cfg.RateLimitCfg,
	logger *zap.Logger,
	settings *Settings,
) (*RateLimiter, error) {
	instance := &RateLimiter{
		mu:       sync.RWMutex{},
		buckets:  make(map[string]*TokenBucket),
		logger:   logger,
		settings: settings,
	}

	script, err := os.ReadFile(rlCfg.RedisLuaScriptPath)
	if err != nil {
		return nil, err
	}

	instance.redisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.Database,
	})
	instance.redisScript = redis.NewScript(string(script))

	return instance, err
}

func (l *RateLimiter) IsAllowed(ctx context.Context, requestType RequestType, identifier string) bool {
	data := fmt.Sprintf("%s-%s", requestType, identifier)
	h := md5.Sum([]byte(data))
	hash := new(big.Int).SetBytes(h[:]).Text(62)

	reqSettings := l.settings.get(requestType)
	if reqSettings == nil {
		l.logger.Fatal("rate limit settings does not exist", zap.String("requestType", string(requestType)))
		return false
	}

	now := time.Now().UnixMicro()
	res, err := l.redisScript.
		Run(ctx, l.redisClient, []string{hash}, reqSettings.MaxTokens, reqSettings.Rate, now, 1).
		Result()
	if err != nil {
		l.logger.Error("run script", zap.Error(err))
		return l.IsAllowedLocal(requestType, defaultTokenSize, identifier)
	}

	allowed, ok := res.(int64)
	if !ok {
		return false
	}

	return allowed == 1
}

func (l *RateLimiter) IsAllowedLocal(requestType RequestType, tokens float64, identifier string) bool {
	bucket := l.get(identifier)
	if bucket == nil {
		reqSettings := l.settings.get(requestType)
		if reqSettings == nil {
			l.logger.Fatal("rate limit settings does not exist", zap.String("requestType", string(requestType)))
			return false
		}

		bucket = l.set(identifier, reqSettings)
	}

	return bucket.IsAllowed(tokens)
}

func (l *RateLimiter) get(identifier string) *TokenBucket {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.buckets[identifier]
}

func (l *RateLimiter) set(identifier string, settings *settings) *TokenBucket {
	l.mu.Lock()
	defer l.mu.Unlock()

	bucket := NewTokenBucket(settings)
	l.buckets[identifier] = bucket

	return bucket
}
