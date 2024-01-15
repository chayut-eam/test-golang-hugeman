package meta

type HealthCheckConfig struct {
	CacheDuration   int64 `mapstructure:"cacheDuration"`
	RefreshInterval int64 `mapstructure:"refreshInterval"`
	InitialDelay    int64 `mapstructure:"initialDelay"`
}
