package config

import (
	"context"
	"net/http"
	"time"

	"github.com/chayut-eam/test-golang-hugeman/handler/event"
	"github.com/chayut-eam/test-golang-hugeman/handler/meta"
	"github.com/chayut-eam/test-golang-hugeman/model"
	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"
)

type APIServer struct {
	server *http.Server
	addr   string
	logger *logrus.Entry
}

func newAPIServer(ApiServerConfig model.APIServerConfig, appInfo model.AppInfo, eventHandler event.Handler, metaHandler meta.MetaHandler, logger *logrus.Entry) App {
	router := gin.Default()

	// Set Gin mode and other configurations
	gin.SetMode(gin.ReleaseMode)

	router = gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/info", metaHandler.Info())
	router.GET("/health", metaHandler.Health())
	router.GET("/data", eventHandler.GetAllHandler)
	router.POST("/create", eventHandler.CreateHandler)
	router.GET("data/:search_value", eventHandler.SearchHandler)
	router.PATCH("update/:id", eventHandler.UpdateHandler)

	app := &APIServer{
		server: &http.Server{
			Addr:    ":" + ApiServerConfig.Port,
			Handler: router,
		},
		addr:   ApiServerConfig.Port,
		logger: logger,
	}

	return app
}

func (this *APIServer) Start() {
	this.logger.Infof("Starting API Server on port : %v ...", this.addr)
	if err := this.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		this.logger.Fatalf("Error starting API Server. Caused by: %v.", err)
	}
}

func (this *APIServer) Stop() {
	this.logger.Info("Stopping API Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := this.server.Shutdown(ctx); err != nil {
		this.logger.Errorf("Error stopping API Server. Caused by: %v.", err)
	}
}
