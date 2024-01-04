package config

const (
	defaultServiceName = "nftique"
	defaultNamespace   = "nftique_namespace"
	defaultHost        = "127.0.0.1"
	defaultPort        = "8080"
	defaultMonitorHost = "127.0.0.1"
	defaultMonitorPort = "8090"
	defaultEnvironment = "dev"

	defaultRedisDB = 0
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
)
