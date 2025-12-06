package main

import (
	"api-gateway/config"
	"api-gateway/internal/router"
	"api-gateway/pkg/logger"

	_ "api-gateway/docs"

	"github.com/gin-gonic/gin"
)

// @title           Taxi Hub API Gateway
// @version         1.0
// @description     API Gateway for Taxi Hub microservices
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    https://www.github.com/enesayyildiz1
// @contact.email  ayyildiz_enes66@hotmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-Key
// @description API Key for admin endpoints

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
