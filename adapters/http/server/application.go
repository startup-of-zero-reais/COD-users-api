package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/routes"
)

type (
	ApplicationInterface interface {
		Start()
		Route()
	}

	Application struct {
		e *echo.Echo
	}
)

func NewApplication() *Application {
	s := echo.New()
	s.Use(middleware.Logger())

	return &Application{
		e: s,
	}
}

func (a *Application) Start() {
	a.Router()
	a.e.Logger.Fatal(a.e.Start(":8080"))
}

func (a *Application) Router() {
	a.e.GET(routes.HealthCheck())

	routes.NewUser(a.e.Group("/users")).Register()
}
