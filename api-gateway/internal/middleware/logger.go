package middleware

import (
	"api-gateway/pkg/logger"
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		var requestBody string
		if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			requestBody = string(bodyBytes)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		username := c.GetString("username")
		userID := c.GetString("user_id")
		role := c.GetString("role")

		c.Next()

		duration := time.Since(startTime)
		statusCode := c.Writer.Status()

		var logLevel logrus.Level
		if statusCode >= 500 {
			logLevel = logrus.ErrorLevel
		} else if statusCode >= 400 {
			logLevel = logrus.WarnLevel
		} else {
			logLevel = logrus.InfoLevel
		}

		fields := logrus.Fields{
			"method":      c.Request.Method,
			"path":        c.Request.URL.Path,
			"query":       c.Request.URL.RawQuery,
			"status":      statusCode,
			"duration_ms": duration.Milliseconds(),
			"ip":          c.ClientIP(),
			"user_agent":  c.Request.UserAgent(),
		}

		if username != "" {
			fields["username"] = username
			fields["user_id"] = userID
			fields["role"] = role
		}

		if logger.Log.Level == logrus.DebugLevel && requestBody != "" {
			fields["request_body"] = maskSensitiveData(requestBody)
		}

		if len(c.Errors) > 0 {
			fields["errors"] = c.Errors.String()
		}

		logger.Log.WithFields(fields).Log(logLevel, "HTTP Request")
	}
}

func maskSensitiveData(data string) string {
	if len(data) > 500 {
		return data[:500] + "... (truncated)"
	}
	return data
}
