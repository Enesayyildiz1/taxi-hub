package main

import (
	"api-gateway/config"
	"api-gateway/internal/router"
	"api-gateway/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.Init()
	cfg := config.Load()
	gin.SetMode(cfg.GinMode)
	r := router.SetupRouter(cfg)

	logger.Log.WithField("port", cfg.Port).Info("API Gateway starting")
	logger.Log.WithField("driver_service", cfg.DriverServiceURL).Info("Driver service URL configured")

	if err := r.Run(":" + cfg.Port); err != nil {
		logger.Log.WithError(err).Fatal("Server failed to start")
	}
}
