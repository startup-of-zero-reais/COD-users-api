package user

import (
	"github.com/labstack/echo/v4"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/controllers/router"
	"net/http"
)

func (u *User) listHandler(c echo.Context) error {
	users := u.Service.List([]string{}, 1, 10)
	return c.JSON(http.StatusOK, users)
}

func (u *User) List() {
	route := router.NewRoute(u.Group)
	route.Register(u.listHandler)
	u.register(route)
}
