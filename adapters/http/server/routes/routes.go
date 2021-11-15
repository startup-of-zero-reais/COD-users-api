package routes

import "github.com/labstack/echo/v4"

type (
	Router interface {
		Register()
	}

	RouterHandler struct {
		Path    string
		Handler func(c echo.Context) error
		Method  string
		Group   *echo.Group
	}
)

func NewRouter(g *echo.Group) *RouterHandler {
	return &RouterHandler{
		Path:   "",
		Method: "GET",
		Group:  g,
	}
}

func (rh *RouterHandler) Route() (string, func(c echo.Context) error) {
	return rh.Path, rh.Handler
}

func (rh *RouterHandler) registerMethod() {
	if rh.IsValidRoute() {
		switch rh.Method {
		case "GET":
			rh.Group.GET(rh.Route())
		case "POST":
			rh.Group.POST(rh.Route())
		case "PUT":
			rh.Group.PUT(rh.Route())
		case "PATCH":
			rh.Group.PATCH(rh.Route())
		case "DELETE":
			rh.Group.DELETE(rh.Route())
		default:
			rh.Group.GET(rh.Route())
		}
	}
}

func (rh *RouterHandler) IsValidRoute() bool {
	acceptedMethods := []string{
		"GET",
		"POST",
		"PUT",
		"DELETE",
		"PATCH",
	}

	if contains(acceptedMethods, rh.Method) && rh.Handler != nil {
		return true
	}

	return false
}

func contains(haystack []string, el string) bool {
	for _, needle := range haystack {
		if el == needle {
			return true
		}
	}

	return false
}
