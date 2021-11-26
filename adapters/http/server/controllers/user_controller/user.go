package user_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/database"
	paginatorsAdapter "github.com/startup-of-zero-reais/COD-users-api/adapters/http/paginators"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/controllers/router"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/middlewares/x_api_key"
	servicesAdapter "github.com/startup-of-zero-reais/COD-users-api/adapters/http/services"
	validatorsAdapter "github.com/startup-of-zero-reais/COD-users-api/adapters/http/validators"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/paginators"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/services"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/validators"
)

type (
	// User é o controller principal da aplicação
	//
	// User é a estrutura do controller de entities.User
	User struct {
		Routes    []*router.Route
		Service   services.UserService
		Validator validators.UserValidator
		Paginator paginators.Pager
		Group     *echo.Group
	}
)

// New é o construtor do controller de User
func New(g *echo.Group, db *database.Database) *User {
	return &User{
		Service:   servicesAdapter.NewUser(db),
		Validator: validatorsAdapter.NewUser(),
		Paginator: paginatorsAdapter.NewPaginator("/users"),
		Group:     g,
	}
}

// Register é o método que implementa o Controller
//
// Register faz com que User seja um controller.
// Este método registra os middlewares e rotas do controller de usuários
func (u *User) Register() {
	u.Middlewares()

	u.List()
	u.Create()
	u.Update()
	u.Delete()

	for _, route := range u.Routes {
		route.RegisterRoutes()
	}
}

// Middlewares registra os middlewares necessários no controller de usuários
func (u *User) Middlewares() {
	apiAuth := x_api_key.NewXApiKey()
	checkMiddleware := (func(h echo.HandlerFunc) echo.HandlerFunc)(apiAuth.CheckApplication())
	keyAuth := (func(h echo.HandlerFunc) echo.HandlerFunc)(apiAuth.KeyAuth())

	u.Group.Use(checkMiddleware, keyAuth)
}

// O register registra as rotas dentro do Controller de User
func (u *User) register(route *router.Route) {
	u.Routes = append(u.Routes, route)
}
