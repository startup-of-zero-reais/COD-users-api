package main

import (
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server"
	"log"
	"os"
)
import "github.com/joho/godotenv"

func init() {
	if l := os.Getenv("APPLICATION_ENV"); l == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}
}

func main() {
	app := server.NewApplication()
	app.Start()
}
