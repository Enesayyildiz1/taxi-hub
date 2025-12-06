// internal/router/router.go
package router

import (
	"api-gateway/config"
	"api-gateway/internal/handler"
	"api-gateway/internal/middleware"
	"api-gateway/internal/proxy"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.ErrorLogger())
	r.Use(middleware.RequestLogger())

	rateLimiter := middleware.NewRateLimiter(cfg.RateLimitRPS, cfg.RateLimitRPS*2)
	r.Use(rateLimiter.Middleware())

	healthHandler := handler.NewHealthHandler()
	authHandler := handler.NewAuthHandler(cfg.JWTSecret)
	driverProxy := proxy.NewDriverProxy(cfg.DriverServiceURL)
	driverHandler := handler.NewDriverHandler(driverProxy)
	adminHandler := handler.NewAdminHandler()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/health", healthHandler.HealthCheck)

	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
	}

	api := r.Group("/api/v1")
	{
		public := api.Group("")
		public.Use(middleware.OptionalJWTAuth(cfg.JWTSecret))
		{
			public.GET("/drivers/nearby", driverHandler.GetNearbyDrivers)
			public.GET("/drivers", driverHandler.GetDrivers)
			public.GET("/drivers/:id", driverHandler.GetDriver)
		}

		protected := api.Group("")
		protected.Use(middleware.JWTAuth(cfg.JWTSecret))
		{
			protected.POST("/drivers", driverHandler.CreateDriver)
			protected.PUT("/drivers/:id", driverHandler.UpdateDriver)
			protected.DELETE("/drivers/:id", driverHandler.DeleteDriver)
		}

		admin := api.Group("/admin")
		admin.Use(middleware.JWTAuth(cfg.JWTSecret))
		admin.Use(middleware.APIKeyAuth(cfg.APIKey))
		{
			admin.GET("/stats", adminHandler.GetStats)
		}
	}

	return r
}
