package router

import "github.com/labstack/echo/v4"

type (
	Method            string
	MiddlewareHandler func(c echo.HandlerFunc) echo.HandlerFunc
	Handler           func(c echo.Context) error

	Controller interface {
		Register()
	}

	Route struct {
		Path        string
		Middlewares map[string][]MiddlewareHandler
		Handler     Handler
		Method      Method
		Group       *echo.Group
	}
)

const (
	GET    = Method("GET")
	POST   = Method("POST")
	PUT    = Method("PUT")
	PATCH  = Method("PATCH")
	DELETE = Method("DELETE")
)

func NewRoute(g *echo.Group) *Route {
	return &Route{
		Path:        "",
		Method:      GET,
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

func (r *Route) Use(middlewares ...MiddlewareHandler) {
	for _, middleware := range middlewares {
		r.Middlewares[r.Path] = append(r.Middlewares[r.Path], middleware)
	}
}

func (r *Route) Register(handler Handler) {
	r.Handler = handler
}

func (r *Route) RegisterRoutes() {
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
	acceptedMethods := []Method{
		GET,
		POST,
		PUT,
		DELETE,
		PATCH,
	}

	if contains(acceptedMethods, r.Method) && r.Handler != nil {
		return true
	}

	return false
}

func contains(haystack []Method, el Method) bool {
	for _, needle := range haystack {
		if el == needle {
			return true
		}
	}

	return false
}
