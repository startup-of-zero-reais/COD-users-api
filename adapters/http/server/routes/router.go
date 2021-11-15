package routes

import "github.com/labstack/echo/v4"

type (
	MiddlewareHandler func(c echo.HandlerFunc) echo.HandlerFunc
	Handler           func(c echo.Context) error

	Router interface {
		Register()
	}

	Route struct {
		Path        string
		Middlewares map[string][]MiddlewareHandler
		Handler     Handler
		Method      string
		Group       *echo.Group
	}
)

func NewRoute(g *echo.Group) *Route {
	return &Route{
		Path:        "",
		Method:      "GET",
		Group:       g,
		Middlewares: map[string][]MiddlewareHandler{},
	}
}

func (r *Route) Route() (string, func(c echo.Context) error) {
	if r.Middlewares[r.Path] != nil && len(r.Middlewares[r.Path]) > 0 {
		for _, middleware := range r.Middlewares[r.Path] {
			parsedMiddleware := (func(c echo.HandlerFunc) echo.HandlerFunc)(middleware)

			r.Group.Use(parsedMiddleware)
		}
		return r.Path, r.Handler
	}
	return r.Path, r.Handler
}

func (r *Route) use(middlewares ...MiddlewareHandler) {
	for _, middleware := range middlewares {
		r.Middlewares[r.Path] = append(r.Middlewares[r.Path], middleware)
	}
}

func (r *Route) register(handler Handler) {
	r.Handler = handler
}

func (r *Route) registerMethod() {
	if r.IsValidRoute() {
		switch r.Method {
		case "GET":
			r.Group.GET(r.Route())
		case "POST":
			r.Group.POST(r.Route())
		case "PUT":
			r.Group.PUT(r.Route())
		case "PATCH":
			r.Group.PATCH(r.Route())
		case "DELETE":
			r.Group.DELETE(r.Route())
		default:
			r.Group.GET(r.Route())
		}
	}
}

func (r *Route) IsValidRoute() bool {
	acceptedMethods := []string{
		"GET",
		"POST",
		"PUT",
		"DELETE",
		"PATCH",
	}

	if contains(acceptedMethods, r.Method) && r.Handler != nil {
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
