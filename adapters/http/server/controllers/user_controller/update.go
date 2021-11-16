package user_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/controllers/router"
	"net/http"
)

func (u *User) updateHandler(c echo.Context) error {
	id := c.Param("id")

	user, validateErr := u.validate(c)
	if validateErr != nil {
		return c.JSON(http.StatusBadRequest, router.ValidationError("erro de validação", validateErr))
	}

	updatedUser, err := u.Service.Update(id, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, router.HttpError(err))
	}

	return c.JSON(http.StatusOK, updatedUser)

}

func (u *User) Update() {
	route := router.NewRoute(u.Group)
	route.Method = router.PUT
	route.Path = "/:id"
	route.Register(u.updateHandler)
	u.register(route)
}
