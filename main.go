package main

import "github.com/startup-of-zero-reais/COD-users-api/adapters/http/server"

func main() {
	app := server.NewApplication()
	app.Start()
}
