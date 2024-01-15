package logger

import (
	"os"
	"runtime"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/chayut-eam/test-golang-hugeman/model"
)

var (
	logger *log.Entry
)

func Init(appInfo model.AppInfo, logConfig model.LoggerConfig) {
	log.SetFormatter(logFormatter(logConfig))
	log.SetLevel(logLevel(logConfig))
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)

	//logger = log.StandardLogger().WithField("application", appInfo.Name)
	logger = log.StandardLogger().WithFields(log.Fields{
		"application": appInfo.Name,
	})
}

func Logger(logstruct map[string]interface{}) *log.Entry {

	logger = logger.Logger.WithFields(log.Fields{})
	if len(logstruct) > 0 {
		logger = logger.WithFields(logstruct)
	}
	return logger
}

func LoggerSystem() *log.Entry {
	return logger
}

func logFormatter(config model.LoggerConfig) *log.JSONFormatter {

	return &log.JSONFormatter{
		TimestampFormat:  time.RFC3339,
		CallerPrettyfier: ShortCallerPrettyfier,

		FieldMap: log.FieldMap{
			log.FieldKeyFile: "caller",
		},
	}

}

func ShortCallerPrettyfier(frame *runtime.Frame) (function string, file string) {

	fileComponents := strings.Split(frame.Func.Name(), ".")
	nameFnc := fileComponents[len(fileComponents)-1]
	return "", nameFnc

}

func logLevel(config model.LoggerConfig) log.Level {
	logLevel, err := log.ParseLevel(config.LogLevel)
	if err != nil {
		panic(err)
	}
	return logLevel
}
