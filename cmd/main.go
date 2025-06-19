package main

import (
	"eskalate-movie-api/cmd/initiator"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found or error loading it: %v", err)
	}

	// Initialize the application
	r := initiator.InitializeApp()

	// Start the server
	r.Run(":8080")
}
