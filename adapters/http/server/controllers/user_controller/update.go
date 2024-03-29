package user_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/controllers/router"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/middlewares/bearer_jwt"
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

// Update registra a rota de atualização de usuário na aplicação
func (u *User) Update() {
	route := router.NewRoute(u.Group)
	route.Method = router.PUT
	route.Path = "/:id"
	route.Use(bearer_jwt.JwtHeaderConfig())
	route.Register(u.updateHandler)
	u.register(route)
}
