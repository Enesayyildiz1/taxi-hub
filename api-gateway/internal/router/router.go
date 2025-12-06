package router

import (
	"api-gateway/config"
	"api-gateway/internal/middleware"
	"api-gateway/internal/proxy"
	"api-gateway/pkg/jwt"
	"api-gateway/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.ErrorLogger())
	r.Use(middleware.RequestLogger())

	rateLimiter := middleware.NewRateLimiter(cfg.RateLimitRPS, cfg.RateLimitRPS*2)
	r.Use(rateLimiter.Middleware())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "api-gateway",
		})
	})

	auth := r.Group("/auth")
	{
		auth.POST("/login", func(c *gin.Context) {
			var req struct {
				Username string `json:"username" binding:"required"`
				Password string `json:"password" binding:"required"`
			}

			if err := c.ShouldBindJSON(&req); err != nil {
				response.Error(c, http.StatusBadRequest, "Invalid request")
				return
			}

			if req.Username != "admin" || req.Password != "password" {
				response.Error(c, http.StatusUnauthorized, "Invalid credentials")
				return
			}

			token, err := jwt.GenerateToken("user-123", req.Username, "admin", cfg.JWTSecret)
			if err != nil {
				response.Error(c, http.StatusInternalServerError, "Failed to generate token")
				return
			}

			response.Success(c, http.StatusOK, "Login successful", gin.H{
				"token": token,
				"type":  "Bearer",
			})
		})
	}

	driverProxy := proxy.NewDriverProxy(cfg.DriverServiceURL)

	api := r.Group("/api/v1")
	{
		public := api.Group("")
		public.Use(middleware.OptionalJWTAuth(cfg.JWTSecret))
		{
			public.GET("/drivers/nearby", driverProxy.Forward)
			public.GET("/drivers", driverProxy.Forward)
			public.GET("/drivers/:id", driverProxy.Forward)
		}

		protected := api.Group("")
		protected.Use(middleware.JWTAuth(cfg.JWTSecret))
		{
			protected.POST("/drivers", driverProxy.Forward)
			protected.PUT("/drivers/:id", driverProxy.Forward)
			protected.DELETE("/drivers/:id", driverProxy.Forward)
		}

		admin := api.Group("/admin")
		admin.Use(middleware.JWTAuth(cfg.JWTSecret))
		admin.Use(middleware.APIKeyAuth(cfg.APIKey))
		{
			admin.GET("/stats", func(c *gin.Context) {
				response.Success(c, http.StatusOK, "Admin stats", gin.H{
					"total_requests": 1000,
					"active_users":   50,
				})
			})
		}
	}

	return r
}
