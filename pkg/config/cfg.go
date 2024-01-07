package config

import (
	"flag"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var envs = []*EnvVar{
	DefaultServiceName,
	DefaultNamespace,
	DefaultEnvironment,

	DefaultHost,
	DefaultPort,
	DefaultMonitorHost,
	DefaultMonitorPort,

	DefaultRedisAddr,
	DefaultRedisDatabase,
	DefaultRedisPassword,

	DefaultRateLimitLuaScriptPath,
	DefaultRateLimitAuthRate,
	DefaultRateLimitAuthMaxTokens,
	DefaultRateLimitCommonRate,
	DefaultRateLimitCommonMaxTokens,
}

type (
	EnvVar struct {
		DefaultValue           interface{}
		Flag, Env, Description string
	}

	MainCfg struct {
		ServiceName string `mapstructure:"service_name"`
		Namespace   string `mapstructure:"namespace"`
		Environment string `mapstructure:"environment"`

		Host        string `mapstructure:"host"`
		Port        string `mapstructure:"port"`
		MonitorHost string `mapstructure:"monitor_host"`
		MonitorPort string `mapstructure:"monitor_port"`
	}

	RedisCfg struct {
		Addr     string `mapstructure:"redis_address"`
		Password string `mapstructure:"redis_password"`
		Database int    `mapstructure:"redis_database"`
	}

	RateLimitCfg struct {
		RedisLuaScriptPath string `mapstructure:"rate_limit_lua_script_path"`

		AuthRate        float64 `mapstructure:"rate_limit_auth_rate"`
		AuthMaxTokens   float64 `mapstructure:"rate_limit_auth_max_tokens"`
		CommonRate      float64 `mapstructure:"rate_limit_common_rate"`
		CommonMaxTokens float64 `mapstructure:"rate_limit_common_max_tokens"`
	}
)

func NewEnvVar(flag, env string, defaultValue interface{}, description string) *EnvVar {
	return &EnvVar{Flag: flag, Env: env, Description: description, DefaultValue: defaultValue}
}

func AddEnvs(customEnvs []*EnvVar) {
	var tmpEnvs []*EnvVar
	tmpEnvs = append(tmpEnvs, customEnvs...)
	for _, defaultEnv := range envs {
		check := true
		for _, customEnv := range customEnvs {
			if customEnv.Flag == defaultEnv.Flag {
				check = false
				break
			}
		}

		if check {
			tmpEnvs = append(tmpEnvs, defaultEnv)
		}
	}

	envs = tmpEnvs
}

func BindConfig() {
	for _, e := range envs {
		switch val := e.DefaultValue.(type) {
		case string:
			flag.String(e.Flag, val, e.Description)
		case int:
			flag.Int(e.Flag, val, e.Description)
		case bool:
			flag.Bool(e.Flag, val, e.Description)
		case uint64:
			flag.Uint64(e.Flag, val, e.Description)
		case time.Duration:
			flag.Duration(e.Flag, val, e.Description)
		default:
			continue
		}
		if e.DefaultValue != nil {
			viper.SetDefault(e.Env, e.DefaultValue)
		}
	}

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(err)
	}

	var err error
	for _, e := range envs {
		err = viper.BindEnv(e.Env)
		if err != nil {
			panic(err)
		}
	}
}

func GetConfig(cfg interface{}, customEnvs []*EnvVar) error {
	AddEnvs(customEnvs)
	BindConfig()
	return viper.Unmarshal(cfg)
}
