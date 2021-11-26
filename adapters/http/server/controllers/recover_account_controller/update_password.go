package recover_account_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/controllers/router"
	"net/http"
)

type (
	// RecoverPassword é a estrutura de redefinição de senha
	RecoverPassword struct {
		NewPassword             string `json:"new_password"`
		NewPasswordConfirmation string `json:"new_password_confirmation"`
	}
)

// O updatePasswordHandler é o manipulador da rota de atualização senha
func (r *RecoverAccount) updatePasswordHandler(c echo.Context) error {
	id := c.QueryParam("token")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "token não encontrado ou inválido",
		})
	}

	token, err := r.Service.GetToken(id)
	if err != nil || token.Token == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "token não encontrado ou inválido",
		})
	}

	err = r.Service.ValidateToken(token.Token)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]string{
			"message": "token inválido",
		})
	}

	passwordRecover := &RecoverPassword{}
	err = c.Bind(passwordRecover)
	if err != nil || passwordRecover.NewPassword == "" || passwordRecover.NewPassword != passwordRecover.NewPasswordConfirmation {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "as senhas não conferem",
		})
	}

	err = r.Service.UpdatePassword(token.Email, passwordRecover.NewPassword)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "senha atualizada com sucesso!",
	})
}

// UpdatePassword é o método que registra a rota de atualização de senha
func (r *RecoverAccount) UpdatePassword() {
	route := router.NewRoute(r.Group)
	route.Method = router.POST
	route.Path = "/update-password"
	route.Register(r.updatePasswordHandler)
	r.register(route)
}
