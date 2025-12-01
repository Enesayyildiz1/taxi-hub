package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI   string
	DBName     string
	ServerPort string
	GinMode    string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}

	return &Config{
		MongoURI:   getEnv("MONGO_URI", "mongodb://localhost:27017"),
		DBName:     getEnv("DB_NAME", "taxihub"),
		ServerPort: getEnv("SERVER_PORT", "8081"),
		GinMode:    getEnv("GIN_MODE", "debug"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
