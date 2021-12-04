package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/database"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/controllers"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/controllers/auth_controller"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/controllers/recover_account_controller"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/controllers/user_controller"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/middlewares/cors"
	"os"
)

type (
	// ApplicationInterface é a 'interface' para rodar a aplicação
	ApplicationInterface interface {
		Start()
		Route()
	}

	// Application é a estrutura do server
	Application struct {
		e  *echo.Echo
		db *database.Database
	}
)

// O makeCodDbConnection cria uma instância de database.Database e conecta a base de dados
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

// NewApplication gera uma nova instância de Application
func NewApplication() *Application {
	s := echo.New()
	s.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339}","id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency_human":"${latency_human}"` + "\n",
	}))

	db, err := makeCodDbConnection()
	if err != nil {
		s.Logger.Fatal(err)
	}

	return &Application{
		e:  s,
		db: db,
	}
}

// Start roda a Application
func (a *Application) Start() {
	a.Router()
	a.e.Logger.Fatal(a.e.Start(":8080"))
}

// Router cria o roteamento da Application
func (a *Application) Router() {
	a.e.GET(controllers.HealthCheck())
	a.e.POST(controllers.GenKey())
	a.e.Use(cors.Middleware())

	auth_controller.New(
		a.e.Group("/auth"),
		a.db,
	).Register()
	user_controller.New(
		a.e.Group("/users"),
		a.db,
	).Register()
	recover_account_controller.New(
		a.e.Group("/recover-account"),
		a.db,
	).Register()
}
