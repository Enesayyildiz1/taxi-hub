package router

import (
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

	return r
}
