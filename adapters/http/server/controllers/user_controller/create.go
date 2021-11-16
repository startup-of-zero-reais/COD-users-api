package user_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/controllers/router"
	"net/http"
)

func (u *User) createHandler(c echo.Context) error {
	user, validateErr := u.validate(c)
	if validateErr != nil {
		return c.JSON(http.StatusBadRequest, router.ValidationError("erro de validação", validateErr))
	}

	createdUser, err := u.Service.Create(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, router.HttpError(err))
	}

	return c.JSON(http.StatusCreated, createdUser)
}

func (u *User) Create() {
	route := router.NewRoute(u.Group)
	route.Method = router.POST
	route.Register(u.createHandler)
	u.register(route)
}
