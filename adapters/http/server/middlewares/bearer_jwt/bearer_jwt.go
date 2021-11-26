package bearer_jwt

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/controllers/router"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/services"
	"github.com/startup-of-zero-reais/COD-users-api/domain/utilities"
)

// JwtHeaderConfig retorna o middleware que valida o jwt de autenticação
func JwtHeaderConfig() router.MiddlewareHandler {
	config := middleware.JWTConfig{
		Claims:     &services.JwtCustomClaims{},
		SigningKey: []byte(utilities.GetEnv("APP_SECRET", "tokn-secre7")),
	}

	return (router.MiddlewareHandler)(middleware.JWTWithConfig(config))
}
