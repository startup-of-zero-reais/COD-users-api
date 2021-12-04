package auth_controller

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/database"
	paginatorsAdapter "github.com/startup-of-zero-reais/COD-users-api/adapters/http/paginators"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/controllers/router"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/middlewares/bearer_jwt"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/middlewares/x_api_key"
	servicesAdapter "github.com/startup-of-zero-reais/COD-users-api/adapters/http/services"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/paginators"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/services"
	"github.com/startup-of-zero-reais/COD-users-api/domain/utilities"
	"net/http"
	"strings"
	"time"
)

type (
	// Auth é a estrutura de rotas de autenticação da API
	Auth struct {
		Routes    []*router.Route
		Service   services.AuthService
		Paginator paginators.Pager
		Group     *echo.Group
	}

	// A authBody é a estrutura de autenticação do corpo da requisição
	authBody struct {
		Email    string `json:"email,omitempty"`
		Password string `json:"password,omitempty"`
	}
)

// New é um construtor do controller de Auth
func New(g *echo.Group, db *database.Database) *Auth {
	return &Auth{
		Service:   servicesAdapter.NewAuth(db),
		Paginator: paginatorsAdapter.NewPaginator("/auth"),
		Group:     g,
	}
}

// Register é o método necessário para que Auth seja um Controller
//
// Register é responsável por executar os registros de rotas
func (a *Auth) Register() {
	a.Login()
	a.Me()

	for _, r := range a.Routes {
		r.RegisterRoutes()
	}
}

// O loginHandler é o método manipulador para executar o login de usuário
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
		Name:      user.Name,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Type:      user.Type,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		StandardClaims: jwt.StandardClaims{
			Subject:   user.ID,
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

// Login é o registro da rota de autenticação na API
func (a *Auth) Login() {
	route := router.NewRoute(a.Group)
	route.Method = router.POST
	route.Register(a.loginHandler)
	a.register(route)
}

// O meHandler é o método para retornar o usuário que está logado baseado no header JWT
func (a *Auth) meHandler(c echo.Context) error {
	tokenString := strings.Split(c.Request().Header.Get("Authorization"), "Bearer ")[1]
	user, err := bearer_jwt.JwtDecode(tokenString)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]string{
			"message": fmt.Sprintf("Unauthorized: %s", err.Error()),
		})
	}

	user.HideSensitiveFields()
	return c.JSON(http.StatusOK, user)
}

// Me é o registro da rota de getMe
func (a *Auth) Me() {
	route := router.NewRoute(a.Group)
	route.Path = "/me"

	apiAuth := x_api_key.NewXApiKey()
	checkMiddleware := (func(h echo.HandlerFunc) echo.HandlerFunc)(apiAuth.CheckApplication())
	keyAuth := (func(h echo.HandlerFunc) echo.HandlerFunc)(apiAuth.KeyAuth())

	route.Use(bearer_jwt.JwtHeaderConfig())
	route.Use(checkMiddleware)
	route.Use(keyAuth)

	route.Register(a.meHandler)
	a.register(route)
}

// O register é o método para registrar as Rotas criadas no Controller de Auth
func (a *Auth) register(route *router.Route) {
	a.Routes = append(a.Routes, route)
}
