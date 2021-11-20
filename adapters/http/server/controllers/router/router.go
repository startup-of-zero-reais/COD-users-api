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

func (r *Route) Route() (string, func(c echo.Context) error, []echo.MiddlewareFunc) {
	return r.Path, r.Handler, r.extractMiddlewares()
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
		path, handler, middlewares := r.Route()

		switch r.Method {
		case "GET":
			r.Group.GET(path, handler, middlewares...)
		case "POST":
			r.Group.POST(path, handler, middlewares...)
		case "PUT":
			r.Group.PUT(path, handler, middlewares...)
		case "PATCH":
			r.Group.PATCH(path, handler, middlewares...)
		case "DELETE":
			r.Group.DELETE(path, handler, middlewares...)
		default:
			r.Group.GET(path, handler, middlewares...)
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

func (r *Route) extractMiddlewares() []echo.MiddlewareFunc {
	var m []echo.MiddlewareFunc
	for _, middleware := range r.Middlewares[r.Path] {
		m = append(m, (func(c echo.HandlerFunc) echo.HandlerFunc)(middleware))
	}

	return m
}

func contains(haystack []Method, el Method) bool {
	for _, needle := range haystack {
		if el == needle {
			return true
		}
	}

	return false
}
