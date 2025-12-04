package router

import (
	"driver-service/internal/handler"
	"driver-service/internal/repository"
	"driver-service/internal/service"
	"driver-service/pkg/database"

	"github.com/gin-gonic/gin"
)

func SetupRouter(mongoDB *database.MongoDB) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "OK",
			"service": "driver-service",
		})
	})

	driverCollection := mongoDB.GetCollection("drivers")
	driverRepo := repository.NewDriverRepository(driverCollection)
	driverService := service.NewDriverService(driverRepo)
	driverHandler := handler.NewDriverHandler(driverService)

	api := r.Group("/api/v1")
	{
		drivers := api.Group("/drivers")
		{
			drivers.POST("", driverHandler.CreateDriver)
			drivers.GET("", driverHandler.GetDrivers)
			drivers.GET("/nearby", driverHandler.GetNearbyDrivers)
			drivers.GET("/:id", driverHandler.GetDriver)
			drivers.PUT("/:id", driverHandler.UpdateDriver)
			drivers.DELETE("/:id", driverHandler.DeleteDriver)
		}
	}

	return r
}
