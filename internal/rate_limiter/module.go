package rate_limiter

import (
	"context"

	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewRateLimiter,
	),
	fx.Provide(func() *redis.Client {
		return nil
	}),
	fx.Provide(
		func() (*Settings, error) {
			settings := NewSettings()

			if err := settings.Set(Auth, 5, 1); err != nil {
				return nil, err
			}

			if err := settings.Set(Common, 5, 1); err != nil {
				return nil, err
			}

			return settings, nil
		},
	),
	fx.Invoke(func(ls fx.Lifecycle) {
		ls.Append(fx.Hook{
			OnStart: func(_ context.Context) error {
				return nil
			},
			OnStop: func(_ context.Context) error {
				return nil
			},
		})
	}),
)
