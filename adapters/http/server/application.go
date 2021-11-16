package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/database"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/routes"
	"os"
)

type (
	ApplicationInterface interface {
		Start()
		Route()
	}

	Application struct {
		e  *echo.Echo
		db *database.Database
	}
)

func makeCodDbConnection() (*database.Database, error) {
	dsn := os.Getenv("MYSQL_COD_DSN")
	db := database.NewDatabase()
	db.Dsn = dsn
	err := db.Connect()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewApplication() *Application {
	s := echo.New()
	s.Use(middleware.Logger())

	db, err := makeCodDbConnection()
	if err != nil {
		s.Logger.Fatal(err)
	}

	return &Application{
		e:  s,
		db: db,
	}
}

func (a *Application) Start() {
	a.Router()
	a.e.Logger.Fatal(a.e.Start(":8080"))
}

func (a *Application) Router() {
	a.e.GET(routes.HealthCheck())

	routes.NewUser(
		a.e.Group("/users"),
		a.db,
	).Register()
}
