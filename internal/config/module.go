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
	fx.Provide(func(cfg *Config) *config.DBCfg {
		return &cfg.DBCfg
	}),
	fx.Provide(func(cfg *Config) *config.ETHCfg {
		return &cfg.ETHCfg
	}),
	fx.Provide(func(cfg *Config) *config.AWSCfg { return &cfg.AWSCfg }),
)
