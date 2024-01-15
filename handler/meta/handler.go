package meta

import (
	_ "embed"
	"net/http"

	"github.com/alexliesenfeld/health"
	"github.com/gin-gonic/gin"
)

type MetaHandler interface {
	Info() gin.HandlerFunc
	Health() gin.HandlerFunc
}

type MetaHandlerImpl struct {
	info               InfoResponse
	healthChecker      health.Checker
	healthCheckOptions health.HandlerOption
}

func NewMetaHandlerImpl(builtInfo string, healthCheckConfig HealthCheckConfig) MetaHandler {
	return &MetaHandlerImpl{
		info:               ParseInfo(builtInfo),
		healthChecker:      NewHealthChecker(healthCheckConfig),
		healthCheckOptions: health.WithResultWriter(&JsonResultWriter{}),
	}
}

func (this *MetaHandlerImpl) Info() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"info": this.info})
	}
}

// func (this *MetaHandlerImpl) Health() gin.HandlerFunc {
// 	return adaptor.HTTPHandlerFunc(health.NewHandler(this.healthChecker, this.healthCheckOptions))
// }

func (this *MetaHandlerImpl) Health() gin.HandlerFunc {
	// Adapt your health check logic here
	return func(c *gin.Context) {
		result := this.healthChecker.Check(c.Request.Context())

		// Assuming you have a specific response structure for health checks
		c.JSON(http.StatusOK, gin.H{"status": result.Status, "details": result.Details})
	}
}
