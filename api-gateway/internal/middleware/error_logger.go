// internal/middleware/error_logger.go
package middleware

import (
	"api-gateway/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ErrorLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Log.WithFields(logrus.Fields{
					"error":  err,
					"path":   c.Request.URL.Path,
					"method": c.Request.Method,
					"ip":     c.ClientIP(),
				}).Error("Panic recovered")

				c.JSON(500, gin.H{
					"error": "Internal server error",
				})
				c.Abort()
			}
		}()

		c.Next()
	}
}
