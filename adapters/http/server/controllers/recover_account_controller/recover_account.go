package recover_account_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/database"
	paginatorsAdapter "github.com/startup-of-zero-reais/COD-users-api/adapters/http/paginators"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/controllers/router"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/middlewares/x_api_key"
	servicesAdapter "github.com/startup-of-zero-reais/COD-users-api/adapters/http/services"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/paginators"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/services"
	"net/http"
)

type (
	// RecoverAccount é a estrutura do Controller de recuperação de conta
	RecoverAccount struct {
		Routes    []*router.Route
		Service   services.RecoverAccountService
		Paginator paginators.Pager
		Group     *echo.Group
	}
)

// New é o construtor do controller de RecoverAccount
func New(g *echo.Group, db *database.Database) *RecoverAccount {
	return &RecoverAccount{
		Service:   servicesAdapter.NewRecoverAccount(db),
		Paginator: paginatorsAdapter.NewPaginator("/recover-account"),
		Group:     g,
	}
}

// O recoverHandler é o manipulador da rota de recuperação de conta
func (r *RecoverAccount) recoverHandler(c echo.Context) error {
	body := struct {
		Email string `json:"email,omitempty"`
	}{}

	err := c.Bind(&body)
	if err != nil {
		return err
	}

	hasBeenSent := r.Service.SendEmail(body.Email)
	if !hasBeenSent {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "E-mail não enviado. Confira a credencial",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "E-mail enviado!",
	})
}

// Recover é o método que registra a rota de recuperação de conta
func (r *RecoverAccount) Recover() {
	route := router.NewRoute(r.Group)
	route.Method = router.POST
	route.Register(r.recoverHandler)
	r.register(route)
}

// Register é o método que implementa o Controller
//
// Register é o método que registra as rotas e middlewares do controller de RecoverAccount
func (r *RecoverAccount) Register() {
	r.Middlewares()

	r.Recover()
	r.UpdatePassword()

	for _, r := range r.Routes {
		r.RegisterRoutes()
	}
}

// Middlewares registra os middlewares no grupo de rotas de RecoverAccount
func (r *RecoverAccount) Middlewares() {
	apiAuth := x_api_key.NewXApiKey()
	checkMiddleware := (func(h echo.HandlerFunc) echo.HandlerFunc)(apiAuth.CheckApplication())
	keyAuth := (func(h echo.HandlerFunc) echo.HandlerFunc)(apiAuth.KeyAuth())

	r.Group.Use(checkMiddleware, keyAuth)
}

// O register registra as rotas criadas no Controller de RecoverAccount
func (r *RecoverAccount) register(route *router.Route) {
	r.Routes = append(r.Routes, route)
}
