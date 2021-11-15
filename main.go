package main

import (
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server"
	"log"
)
import "github.com/joho/godotenv"

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	app := server.NewApplication()
	app.Start()
}
