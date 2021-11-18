package user_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/database"
	paginatorsAdapter "github.com/startup-of-zero-reais/COD-users-api/adapters/http/paginators"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/controllers/router"
	servicesAdapter "github.com/startup-of-zero-reais/COD-users-api/adapters/http/services"
	validatorsAdapter "github.com/startup-of-zero-reais/COD-users-api/adapters/http/validators"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/paginators"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/services"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/validators"
)

type (
	User struct {
		Routes    []*router.Route
		Service   services.UserService
		Validator validators.UserValidator
		Paginator paginators.Pager
		Group     *echo.Group
	}
)

func New(g *echo.Group, db *database.Database) *User {
	return &User{
		Service:   servicesAdapter.NewUser(db),
		Validator: validatorsAdapter.NewUser(),
		Paginator: paginatorsAdapter.NewPaginator("/users"),
		Group:     g,
	}
}

func (u *User) Register() {
	u.List()
	u.Create()
	u.Update()
	u.Delete()

	for _, route := range u.Routes {
		route.RegisterRoutes()
	}
}

func (u *User) register(route *router.Route) {
	u.Routes = append(u.Routes, route)
}