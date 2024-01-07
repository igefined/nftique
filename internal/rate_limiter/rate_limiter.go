package rate_limiter

import (
	"sync"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

const (
	Auth   RequestType = "auth"
	Common RequestType = "common"

	Undefined RequestType = "undefined"
)

type RequestType string

type RateLimiter struct {
	mu      sync.RWMutex
	buckets map[string]*TokenBucket

	logger   *zap.Logger
	settings *Settings
	redis    *redis.Client
}

func NewRateLimiter(logger *zap.Logger, settings *Settings, redis *redis.Client) *RateLimiter {
	return &RateLimiter{
		mu:       sync.RWMutex{},
		buckets:  make(map[string]*TokenBucket),
		redis:    redis,
		logger:   logger,
		settings: settings,
	}
}

func (l *RateLimiter) IsAllowed(requestType RequestType, tokens float64, identifier string) bool {
	reqSettings := l.settings.get(requestType)
	if reqSettings == nil {
		l.logger.Fatal("rate limit settings does not exist", zap.String("requestType", string(requestType)))
		return false
	}

	bucket := l.get(identifier)
	if bucket == nil {
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
