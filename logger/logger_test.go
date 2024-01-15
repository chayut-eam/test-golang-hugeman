package logger_test

import (
	"testing"

	logs "github.com/chayut-eam/test-golang-hugeman/logger"
	"github.com/chayut-eam/test-golang-hugeman/model"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	config := model.AppConfig{
		AppInfo: model.AppInfo{
			Name: "test",
		},
		LoggerConfig: model.LoggerConfig{
			LogLevel: "info",
		},
	}
	logs.Init(config.AppInfo, config.LoggerConfig)
	logs.LoggerSystem()

	// log lavel error
	config = model.AppConfig{
		AppInfo: model.AppInfo{
			Name: "test",
		},
		LoggerConfig: model.LoggerConfig{
			LogLevel: "test",
		},
	}
	assert.Panics(t, func() {
		logs.Init(config.AppInfo, config.LoggerConfig)
	})

	data := make(map[string]interface{})
	data["test"] = 46546
	do := logs.Logger(data)
	do.Info("test")

}
