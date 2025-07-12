package server

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joseCarlosAndrade/rinha-backend-2025-jc/internal/core/config"
	"github.com/joseCarlosAndrade/rinha-backend-2025-jc/internal/core/domain/port"
	"go.uber.org/zap"
)

type Controller struct {
	Service port.Service
	Gateway *gin.Engine
}


func NewController(serv port.Service) *Controller {
	gateway := gin.New()
	gateway.Use(
		ginZapMiddleware(),
		gin.Recovery(),
	)

	controller := &Controller{
		Service: serv,
		Gateway: gateway,
	}

	controller.Gateway.POST("payments", controller.payments)
	controller.Gateway.GET("payments-summary", controller.paymentsSummary)

	// todo: add swagger

	if !config.App.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	return controller
}


// middleware to log requests when debug is true
func ginZapMiddleware() gin.HandlerFunc {
	return func (c *gin.Context) {

		start := time.Now()
		
		c.Next()

		latency := time.Since(start)
		method := c.Request.Method
		statusCode := c.Writer.Status()
		path := c.Request.URL.Path
		
		if config.App.Debug {
			zap.L().Debug("HTTP",
				zap.String("method", method),
				zap.String("path", path),
				zap.Int("status", statusCode),
				zap.Duration("latency", latency),
			)
		}
	}
}

