package redis

import (
	"context"
	"time"

	"github.com/igefined/nftique/pkg/config"

	"github.com/redis/go-redis/v9"
)

type Redis interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
}

type (
	Client struct {
		cfg *config.RedisCfg

		redis *redis.Client
	}

	Opts struct {
		Expiration time.Duration
	}
)

func New(cfg *config.RedisCfg) *Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.Database,
	})

	return &Client{cfg: cfg, redis: redisClient}
}

func (c *Client) Get(ctx context.Context, key string) (string, error) {
	cmd := c.redis.Get(ctx, key)

	return cmd.Result()
}

func (c *Client) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.redis.Set(ctx, key, value, expiration).Err()
}
