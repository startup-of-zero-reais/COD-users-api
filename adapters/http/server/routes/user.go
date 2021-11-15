package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/services"
	"log"
	"net/http"
)

type (
	User struct {
		Routes  []*RouterHandler
		Service services.UserService
		Group   *echo.Group
	}
)

func NewUser(g *echo.Group) *User {
	return &User{
		Service: nil,
		Group:   g,
	}
}

func (u *User) List() {
	log.Println("registering list user route...")

	listRoute := NewRouter(u.Group)
	listRoute.Handler = func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "List Users ok!"})
	}

	u.Routes = append(u.Routes, listRoute)
}

func (u *User) Register() {
	log.Println("registering user routes...")

	u.List()

	for _, route := range u.Routes {
		route.registerMethod()
	}

	log.Println("done! user routes registered")
}
