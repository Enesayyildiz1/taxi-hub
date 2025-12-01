package main

import (
	"driver-service/config"
	"driver-service/internal/router"
	"driver-service/pkg/database"
	"log"
)

func main() {
	cfg := config.Load()
	mongoDB, err := database.ConnectMongo(cfg.MongoURI, cfg.DBName)

	if err != nil {
		log.Fatal("MongoDB connection failed:", err)
	}
	defer mongoDB.Close()

	r := router.SetupRouter(mongoDB)

	log.Printf("Server started on %s", cfg.ServerPort)
	r.Run(":" + cfg.ServerPort)
}
