package user_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/controllers/router"
	"net/http"
	"net/url"
	"strings"
)

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

	paginated := u.Paginator.Paginate(users)

	return c.JSON(http.StatusOK, paginated)
}

func (u *User) List() {
	route := router.NewRoute(u.Group)
	route.Register(u.listHandler)
	u.register(route)
}