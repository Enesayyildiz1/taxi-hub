package main

import (
	"api-gateway/config"
	"api-gateway/internal/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	gin.SetMode(cfg.GinMode)
	r := router.SetupRouter(cfg)

	log.Printf("API Gateway running on port %s", cfg.Port)
	log.Printf("Forwarding to Driver Service: %s", cfg.DriverServiceURL)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
