package model

import "github.com/chayut-eam/test-golang-hugeman/handler/meta"

type AppConfig struct {
	AppInfo           AppInfo                `mapstructure:"app"`
	APIServerConfig   APIServerConfig        `mapstructure:"apiServer"`
	LoggerConfig      LoggerConfig           `mapstructure:"logging"`
	HealthCheckConfig meta.HealthCheckConfig `mapstructure:"healthCheck"`
}

type AppInfo struct {
	Name string `mapstructure:"name"`
}

type APIServerConfig struct {
	Name         string `mapstructure:"name"`
	Port         string `mapstructure:"port"`
	ReadTimeout  int64  `mapstructure:"readTimeout"`
	WriteTimeout int64  `mapstructure:"writeTimeout"`
	IdleTimeout  int64  `mapstructure:"idleTimeout"`
}

type LoggerConfig struct {
	LogLevel string `mapstructure:"level"`
}
