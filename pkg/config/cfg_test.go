package config

import (
	"os"
	"testing"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestNewEnvVar(t *testing.T) {
	flag := "test_flag"
	env := "test_env"
	description := "test description"

	t.Run("string default", func(t *testing.T) {
		defaultValue := "test_default_value"
		envVar := NewEnvVar(flag, env, defaultValue, description)
		assert.NotNil(t, envVar)
		switch envVar.DefaultValue.(type) {
		case string:
			return
		default:
			t.Fatalf("default value should be string")
		}
	})

	t.Run("int default", func(t *testing.T) {
		defaultValue := 1
		envVar := NewEnvVar(flag, env, defaultValue, description)
		assert.NotNil(t, envVar)
		switch envVar.DefaultValue.(type) {
		case int:
			return
		default:
			t.Fatalf("default value should be string")
		}
	})
}

func TestAddEnvs(t *testing.T) {
	customEnvs := []*EnvVar{
		{
			Flag:         "custom_flag1",
			Env:          "CUSTOM_ENV_1",
			DefaultValue: "default_value_1",
			Description:  "Custom environment variable 1",
		},
		{
			Flag:         "custom_flag2",
			Env:          "CUSTOM_ENV_2",
			DefaultValue: "default_value_2",
			Description:  "Custom environment variable 2",
		},
	}

	AddEnvs(customEnvs)

	for _, customEnv := range customEnvs {
		found := false
		for _, env := range envs {
			if customEnv.Flag == env.Flag {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected custom environment variable %s not found in envs", customEnv.Flag)
		}
	}

	for i, env := range envs {
		for j := i + 1; j < len(envs); j++ {
			if env.Flag == envs[j].Flag {
				t.Errorf("Duplicate environment variable %s found in envs", env.Flag)
			}
		}
	}
}

func TestBindConfig(t *testing.T) {
	oldArgs := os.Args
	t.Cleanup(func() {
		os.Args = oldArgs
		envs = nil
	})

	testArgs := []string{
		"--host=test_host",
		"--port=8080",
		"--monitor_host=test_monitor_host",
		"--custom_flag=test_value",
	}

	os.Args = append([]string{"cmd"}, testArgs...)

	pflag.String("host", "", "nftique host")
	pflag.String("port", "", "nftique port")
	pflag.String("monitor_host", "", "nftique monitor port")
	pflag.String("custom_flag", "", "Custom flag")

	BindConfig()

	assert.Equal(t, "test_host", viper.GetString("HOST"))
	assert.Equal(t, "8080", viper.GetString("PORT"))
	assert.Equal(t, "test_monitor_host", viper.GetString("MONITOR_HOST"))
	assert.Equal(t, "test_value", viper.GetString("CUSTOM_FLAG"))
}

func TestGetConfig(t *testing.T) {
	type C struct {
		MainCfg `mapstructure:",squash"`
		Test    string
	}

	testVars := []*EnvVar{
		NewEnvVar("test", "TEST", "test", "custom test env"),
	}

	var expectedPort = "9090"
	_ = os.Setenv("MONITOR_PORT", expectedPort)

	var c C
	err := GetConfig(&c, testVars)
	assert.NoError(t, err)

	assert.Equal(t, c.MonitorPort, expectedPort)
	assert.Equal(t, c.Test, "test")
}
