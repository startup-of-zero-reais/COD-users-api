package user_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/controllers/router"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/middlewares/bearer_jwt"
	"github.com/startup-of-zero-reais/COD-users-api/domain/entities"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/resources"
	"net/http"
	"net/url"
	"strings"
)

// MaxUsrIds é o número máximo de IDS que pode ser procurado simultaneamente na base
const MaxUsrIds = 128

func (u *User) extractSearch(values url.Values) []string {
	if values.Get("ids") != "" {
		ids := strings.Split(values.Get("ids"), ",")
		if len(ids) > MaxUsrIds {
			return ids[:MaxUsrIds]
		}

		return ids
	}

	return []string{}
}

func (u *User) listHandler(c echo.Context) error {
	page, perPage := u.Paginator.GetPagination(c)

	users, total := u.Service.List(
		u.extractSearch(c.QueryParams()),
		page,
		perPage,
	)

	u.Paginator.Total = uint(total)

	paginated := u.Paginator.Paginate(
		u.GetCollection(users),
	)

	return c.JSON(http.StatusOK, paginated)
}

// GetCollection é o método que implementa o recurso de coleção para User
func (u *User) GetCollection(users []entities.User) []interface{ resources.Resource } {
	var _users []interface{ resources.Resource }
	for _, user := range users {
		_users = append(_users, &user)
	}

	return _users
}

// List registra a rota de listagem de usuários na API
func (u *User) List() {
	route := router.NewRoute(u.Group)
	route.Use(bearer_jwt.JwtHeaderConfig())
	route.Register(u.listHandler)
	u.register(route)
}
