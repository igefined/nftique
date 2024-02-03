package config

import (
	"time"
)

const (
	localhost = "127.0.0.1"

	defaultServiceName = "nftique"
	defaultNamespace   = "nftique_namespace"
	defaultHost        = "127.0.0.1"
	defaultPort        = "8080"
	defaultMonitorHost = localhost
	defaultMonitorPort = "8090"
	defaultEnvironment = "dev"

	defaultRedisDB       = 0
	defaultLuaScriptPath = "rate_limit.lua"

	defaultRate      = 1
	defaultMaxTokens = 60

	defaultDBURL = "postgres://postgres:postgres@localhost:5432/nftique?sslmode=disable"

	defaultChainRPC = "http://127.0.0.1:8567"
)

var (
	DefaultServiceName = NewEnvVar(
		"service_name",
		"SERVICE_NAME",
		defaultServiceName,
		"Service name",
	)

	DefaultNamespace = NewEnvVar(
		"namespace",
		"NAMESPACE",
		defaultNamespace,
		"Service namespace",
	)

	DefaultHost = NewEnvVar(
		"host",
		"HOST",
		defaultHost,
		"Host. Related env var",
	)

	DefaultPort = NewEnvVar(
		"port",
		"PORT",
		defaultPort,
		"Port. Related env var",
	)

	DefaultMonitorHost = NewEnvVar(
		"monitor_host",
		"MONITOR_HOST",
		defaultMonitorHost,
		"Monitor host",
	)

	DefaultMonitorPort = NewEnvVar(
		"monitor_port",
		"MONITOR_PORT",
		defaultMonitorPort,
		"Monitor port",
	)

	DefaultEnvironment = NewEnvVar(
		"environment",
		"ENVIRONMENT",
		defaultEnvironment,
		"Deployment environment",
	)

	DefaultRedisAddr = NewEnvVar(
		"redis_address",
		"REDIS_ADDRESS",
		defaultHost,
		"Redis address",
	)

	DefaultRedisPassword = NewEnvVar(
		"redis_password",
		"REDIS_PASSWORD",
		"",
		"Redis address",
	)

	DefaultRedisDatabase = NewEnvVar(
		"redis_database",
		"REDIS_DATABASE",
		defaultRedisDB,
		"Redis database name",
	)

	DefaultRateLimitLuaScriptPath = NewEnvVar(
		"rate_limit_lua_script_path",
		"RATE_LIMIT_LUA_SCRIPT_PATH",
		defaultLuaScriptPath,
		"Path for redis lua script for rate limiting",
	)

	DefaultRateLimitAuthRate = NewEnvVar(
		"rate_limit_auth_rate",
		"RATE_LIMIT_AUTH_RATE",
		defaultRate,
		"Rate limit auth rate var",
	)

	DefaultRateLimitAuthMaxTokens = NewEnvVar(
		"rate_limit_auth_max_tokens",
		"RATE_LIMIT_AUTH_MAX_TOKENS",
		defaultMaxTokens,
		"Rate limit auth max tokens var",
	)

	DefaultRateLimitCommonRate = NewEnvVar(
		"rate_limit_common_rate",
		"RATE_LIMIT_COMMON_RATE",
		defaultRate,
		"Rate limit common rate var",
	)

	DefaultRateLimitCommonMaxTokens = NewEnvVar(
		"rate_limit_common_max_tokens",
		"RATE_LIMIT_COMMON_MAX_TOKENS",
		defaultMaxTokens,
		"Rate limit common max tokens var",
	)

	DefaultRateLimitNFTsRate = NewEnvVar(
		"rate_limit_nfts_rate",
		"RATE_LIMIT_NFTS_RATE",
		defaultRate,
		"Rate limit NFTs rate var",
	)

	DefaultRateLimitNFTsMaxTokens = NewEnvVar(
		"rate_limit_nfts_max_tokens",
		"RATE_LIMIT_NFTS_MAX_TOKENS",
		defaultMaxTokens,
		"Rate limit NFTs max tokens var",
	)

	DefaultDBURL = NewEnvVar(
		"db_url",
		"DB_URL",
		defaultDBURL,
		"Database url",
	)

	DefaultDBAutoCreate = NewEnvVar(
		"db_auto_create_database",
		"DB_AUTO_CREATE_DATABASE",
		true,
		"The bool variable means that the DB will be created automatically or not",
	)

	DefaultMigrationPath = NewEnvVar(
		"db_migrations_path",
		"DB_MIGRATIONS_PATH",
		"internal/migrations/files",
		"Database migrations path",
	)

	DefaultJWTSecretKey = NewEnvVar(
		"jwt_secret_key",
		"JWT_SECRET_KEY",
		"secret",
		"JWT secret key",
	)

	DefaultExpirationTime = NewEnvVar(
		"jwt_expiration_minutes",
		"JWT_EXPIRATION_MINUTES",
		time.Minute*60,
		"JWT expiration time in minutes",
	)

	DefaultChainRPCtUrl = NewEnvVar(
		"chain_rpc_url",
		"CHAIN_RPC_URL",
		defaultChainRPC,
		"Blockchain rpc url",
	)

	DefaultContractorAddr = NewEnvVar(
		"contractor_address",
		"CONTRACT_ADDRESS",
		"",
		"Blockchain contract address",
	)

	DefaultOperatorAddr = NewEnvVar(
		"operator_address",
		"OPERATOR_ADDRESS",
		"",
		"Blockchain operator address",
	)

	DefaultOperatorPrivateKey = NewEnvVar(
		"operator_private_key",
		"OPERATOR_PRIVATE_KEY",
		"",
		"Blockchain operator private key",
	)
)
