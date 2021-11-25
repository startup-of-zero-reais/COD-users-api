package auth_controller

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/database"
	paginatorsAdapter "github.com/startup-of-zero-reais/COD-users-api/adapters/http/paginators"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/controllers/router"
	servicesAdapter "github.com/startup-of-zero-reais/COD-users-api/adapters/http/services"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/paginators"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/services"
	"github.com/startup-of-zero-reais/COD-users-api/domain/utilities"
	"net/http"
	"time"
)

type (
	Auth struct {
		Routes    []*router.Route
		Service   services.AuthService
		Paginator paginators.Pager
		Group     *echo.Group
	}

	authBody struct {
		Email    string `json:"email,omitempty"`
		Password string `json:"password,omitempty"`
	}
)

func New(g *echo.Group, db *database.Database) *Auth {
	return &Auth{
		Service:   servicesAdapter.NewAuth(db),
		Paginator: paginatorsAdapter.NewPaginator("/auth"),
		Group:     g,
	}
}

func (a *Auth) Register() {
	a.Login()

	for _, r := range a.Routes {
		r.RegisterRoutes()
	}
}

func (a *Auth) loginHandler(c echo.Context) error {
	b := authBody{}

	err := c.Bind(&b)
	if err != nil {
		return err
	}

	user := a.Service.Get(b.Email)

	if user == nil {
		return echo.ErrUnauthorized
	}

	if !user.IsValidPassword(b.Password) {
		return echo.ErrUnauthorized
	}

	claims := &servicesAdapter.JwtCustomClaims{
		Name:  user.Name,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := []byte(utilities.GetEnv("APP_SECRET", "tokn-secre7"))
	t, err := token.SignedString(secret)
	if err != nil {
		return err
	}

	user.HideSensitiveFields()

	return c.JSON(http.StatusOK, echo.Map{
		"user":  user,
		"token": t,
	})
}

func (a *Auth) Login() {
	route := router.NewRoute(a.Group)
	route.Method = router.POST
	route.Register(a.loginHandler)
	a.register(route)
}

func (a *Auth) register(route *router.Route) {
	a.Routes = append(a.Routes, route)
}
