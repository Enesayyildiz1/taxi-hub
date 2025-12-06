package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		log.Printf("[%s] %s %s",
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
		)

		c.Next()

		duration := time.Since(start)
		statusCode := c.Writer.Status()

		log.Printf("[%s] %s - %d (%v)",
			c.Request.Method,
			c.Request.URL.Path,
			statusCode,
			duration,
		)
	}
}
