package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/pkg/logger"
)

func main() {
  env := os.Getenv("APP_ENV")
  if env == "" {
    log.Println("Environment not set, launching on DEV")
    env = "DEV"
  }

	log.Printf("Launching backend API for %v environment...", env)

  log.Println("Config: Loading config...")
  config, err := config.LoadConfig(".env", env)
  if err != nil {
    log.Fatalf("Config: %v", err)
  }
  log.Println("Config: Config loaded !")
  fmt.Println(config)

  log.Println("Logger: Init logger...")
	logger := logger.NewLogger()
	logger.InitLogger()
  log.Println("Logger: Initialized")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World !")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
