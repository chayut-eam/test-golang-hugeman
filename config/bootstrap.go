package config

import (
	"bytes"
	_ "embed"
	"os"
	"os/signal"
	"strings"

	"github.com/chayut-eam/test-golang-hugeman/handler/event"
	"github.com/chayut-eam/test-golang-hugeman/handler/meta"
	log "github.com/chayut-eam/test-golang-hugeman/logger"
	"github.com/chayut-eam/test-golang-hugeman/model"
	"github.com/chayut-eam/test-golang-hugeman/repository"
	"github.com/chayut-eam/test-golang-hugeman/service"
	"github.com/chayut-eam/test-golang-hugeman/validation"

	"github.com/spf13/viper"
)

var (
	//go:embed config.yaml
	conf []byte

	//go:embed info.properties
	builtInfo string
)

type App interface {
	Start()
	Stop()
}

func Bootstrap() App {

	// loading config
	config, err := loadConfig()
	if err != nil {
		log.LoggerSystem().Panicf("Error loading configuration file. Caused by %v", err)
	}

	log.Init(config.AppInfo, config.LoggerConfig)
	logger := log.LoggerSystem()

	logger.Info("Bootstrapping application...")

	// repository
	repository := repository.NewRepository()

	// service
	Service := service.NewService(repository)

	// handler
	eventHandler := event.NewHelloHandlerImpl(Service, logger)
	metaHandler := meta.NewMetaHandlerImpl(builtInfo, config.HealthCheckConfig)

	// app
	apiServer := newAPIServer(config.APIServerConfig, config.AppInfo, eventHandler, metaHandler, logger)

	addShutdownHook(apiServer)

	validation.Init()

	return apiServer
}

func loadConfig() (*model.AppConfig, error) {
	viper.SetConfigType("yaml")
	//viper.SetEnvPrefix("EMAIL")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// read config
	if err := viper.ReadConfig(bytes.NewBuffer(conf)); err != nil {
		return nil, err
	}

	appConfig := &model.AppConfig{}
	if err := viper.Unmarshal(appConfig); err != nil {
		return nil, err
	}

	return appConfig, nil
}

func addShutdownHook(app App) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		_ = <-c
		app.Stop()
	}()
}

func Teardown() {}
