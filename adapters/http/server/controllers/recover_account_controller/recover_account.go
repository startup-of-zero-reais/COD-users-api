package recover_account_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/database"
	paginatorsAdapter "github.com/startup-of-zero-reais/COD-users-api/adapters/http/paginators"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/controllers/router"
	servicesAdapter "github.com/startup-of-zero-reais/COD-users-api/adapters/http/services"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/paginators"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/services"
	"net/http"
)

type (
	RecoverAccount struct {
		Routes    []*router.Route
		Service   services.RecoverAccountService
		Paginator paginators.Pager
		Group     *echo.Group
	}
)

func New(g *echo.Group, db *database.Database) *RecoverAccount {
	return &RecoverAccount{
		Service:   servicesAdapter.NewRecoverAccount(db),
		Paginator: paginatorsAdapter.NewPaginator("/recover-account"),
		Group:     g,
	}
}

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

func (r *RecoverAccount) Recover() {
	route := router.NewRoute(r.Group)
	route.Method = router.POST
	route.Register(r.recoverHandler)
	r.register(route)
}

func (r *RecoverAccount) Register() {
	r.Recover()

	for _, r := range r.Routes {
		r.RegisterRoutes()
	}
}

func (r *RecoverAccount) register(route *router.Route) {
	r.Routes = append(r.Routes, route)
}
