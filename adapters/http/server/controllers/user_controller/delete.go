package user_controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/controllers/router"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/middlewares/bearer_jwt"
	"net/http"
)

func (u *User) deleteHandler(c echo.Context) error {
	id := c.Param("id")

	deleted := u.Service.Delete(id)

	if !deleted {
		return c.JSON(
			http.StatusBadRequest,
			router.HttpError(
				errors.New("usuário não encontrado"),
			),
		)
	}

	return c.JSON(http.StatusNoContent, nil)
}

// Delete registra uma rota para deletar contas de usuário
func (u *User) Delete() {
	route := router.NewRoute(u.Group)
	route.Method = router.DELETE
	route.Path = "/:id"
	route.Use(bearer_jwt.JwtHeaderConfig())
	route.Register(u.deleteHandler)
	u.register(route)
}
