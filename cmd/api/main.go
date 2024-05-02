package main

import (
	"fmt"
	"log"
	"net/http"

    "github.com/esgi-challenge/backend/pkg/logger"
)

func main() {
	log.Println("Launching backend API...")

	logger := logger.NewLogger()
	logger.InitLogger()

	logger.Warn("test")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World !")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
