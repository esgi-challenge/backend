package main

import (
	"log"
	"os"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/server"
	"github.com/esgi-challenge/backend/pkg/logger"
)

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

	s := server.NewServer(config, logger)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
