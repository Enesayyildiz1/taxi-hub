package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port             string
	DriverServiceURL string
	JWTSecret        string
	APIKey           string
	RateLimitRPS     int
	GinMode          string
	LogLevel         string
	LogFormat        string
	EnableFileLog    string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env file not found, using environment variables")
	}

	return &Config{
		Port:             getEnv("PORT", "8080"),
		DriverServiceURL: getEnv("DRIVER_SERVICE_URL", "http://localhost:8081"),
		JWTSecret:        getEnv("JWT_SECRET", "secret-key"),
		APIKey:           getEnv("API_KEY", ""),
		RateLimitRPS:     getEnvInt("RATE_LIMIT_RPS", 10),
		GinMode:          getEnv("GIN_MODE", "debug"),
		LogLevel:         getEnv("LOG_LEVEL", "info"),
		LogFormat:        getEnv("LOG_FORMAT", "json"),
		EnableFileLog:    getEnv("ENABLE_FILE_LOG", "true"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		var intValue int
		if _, err := fmt.Sscanf(value, "%d", &intValue); err == nil {
			return intValue
		}
	}
	return defaultValue
}
