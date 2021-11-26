package router

import "github.com/labstack/echo/v4"

type (
	Method            string
	MiddlewareHandler func(c echo.HandlerFunc) echo.HandlerFunc
	Handler           func(c echo.Context) error

	Controller interface {
		Register()
	}

	// Route é a estrutura de rotas
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

// NewRoute cria uma nota Route
func NewRoute(g *echo.Group) *Route {
	return &Route{
		Path:        "",
		Method:      GET,
		Group:       g,
		Middlewares: map[string][]MiddlewareHandler{},
	}
}

// Route é o método responsável por retornar o caminho, o manipulador e
// caso hajam, os middlewares daquela rota específica
func (r *Route) Route() (string, func(c echo.Context) error, []echo.MiddlewareFunc) {
	return r.Path, r.Handler, r.extractMiddlewares()
}

// Use registra novos middlewares na rota
func (r *Route) Use(middlewares ...MiddlewareHandler) {
	for _, middleware := range middlewares {
		r.Middlewares[r.Path] = append(r.Middlewares[r.Path], middleware)
	}
}

// Register é o método que registra o 'handler' ou manipulador da rota
func (r *Route) Register(handler Handler) {
	r.Handler = handler
}

// RegisterRoutes é o método que verifica se Route é válido.
// Se for valido registra a rota no Group baseado no método HTTP
// Ex.: GET, POST, PUT, DELETE, PATCH
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

// IsValidRoute verifica se Route é uma rota válida, se possui um método válido e
// também, se há um handler para a rota
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

// Este método de extractMiddlewares é apenas para retornar um slice de echo.MiddlewareFunc
// a partir dos Middlewares de Route
func (r *Route) extractMiddlewares() []echo.MiddlewareFunc {
	var m []echo.MiddlewareFunc
	for _, middleware := range r.Middlewares[r.Path] {
		m = append(m, (func(c echo.HandlerFunc) echo.HandlerFunc)(middleware))
	}

	return m
}

// O método contains é para checar se o método de el existe em um slice de Method
func contains(haystack []Method, el Method) bool {
	for _, needle := range haystack {
		if el == needle {
			return true
		}
	}

	return false
}
