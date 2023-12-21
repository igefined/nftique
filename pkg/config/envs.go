package config

const (
	defaultHost        = "127.0.0.1"
	defaultPort        = "8080"
	defaultMonitorHost = "127.0.0.1"
	defaultMonitorPort = "8090"
)

var (
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
)
