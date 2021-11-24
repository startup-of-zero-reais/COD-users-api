package bearer_jwt

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/controllers/auth_controller"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/controllers/router"
	"github.com/startup-of-zero-reais/COD-users-api/domain/utilities"
)

func JwtHeaderConfig() router.MiddlewareHandler {
	config := middleware.JWTConfig{
		Claims:     &auth_controller.JwtCustomClaims{},
		SigningKey: []byte(utilities.GetEnv("APP_SECRET", "tokn-secre7")),
	}

	return (router.MiddlewareHandler)(middleware.JWTWithConfig(config))
}
