package config

import (
	"sync"

	"github.com/igefined/nftique/pkg/config"
)

type Config struct {
	sync.RWMutex
	config.MainCfg      `mapstructure:",squash"`
	config.RedisCfg     `mapstructure:",squash"`
	config.RateLimitCfg `mapstructure:",squash"`
	config.DBCfg        `mapstructure:",squash"`
	config.ETHCfg       `mapstructure:",squash"`
	config.AWSCfg       `mapstructure:",squash"`
}

func New() (*Config, error) {
	c := Config{}
	if err := config.GetConfig(&c, []*config.EnvVar{}); err != nil {
		return nil, err
	}

	return &c, nil
}
