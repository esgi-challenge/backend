package main

import (
	"log"
	"os"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/server"
	"github.com/esgi-challenge/backend/pkg/database"
	"github.com/esgi-challenge/backend/pkg/logger"
)

// @title			Backend
// @description	Backend written in Go for the S2 ESGI Challenge
// @BasePath		/api
func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		log.Println("Environment not set, launching on LOCAL, a .env file is needed then.")
		env = "LOCAL"
	}

	log.Printf("Launching backend API for %v environment...", env)

	log.Println("Config: Loading config...")
	config, err := config.LoadConfig(".env", env)
	if err != nil {
		log.Fatalf("Config: %v", err)
	}
	log.Println("Config: Config loaded !")

	log.Println("Logger: Init logger...")
	logger := logger.NewLogger()
	logger.InitLogger()
	log.Println("Logger: Initialized")

	logger.Info("Database: Init connection")
	psqlDB, err := database.NewPostgresClient(config)
	if err != nil {
		logger.Fatalf("Database: %s", err)
	}
	logger.Info("Database: Postgres connected")

	s := server.NewServer(config, psqlDB, logger)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
