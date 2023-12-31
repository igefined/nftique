package config

import (
	"github.com/igefined/nftique/pkg/config"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(New),
	fx.Provide(func(cfg *Config) *config.RedisCfg {
		return &cfg.RedisCfg
	}),
	fx.Provide(func(cfg *Config) *config.RateLimitCfg {
		return &cfg.RateLimitCfg
	}),
)
