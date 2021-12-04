package cors

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/startup-of-zero-reais/COD-users-api/domain/utilities"
	"strings"
)

func Middleware() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: strings.Split(utilities.GetEnv("ACCEPT_ORIGINS", "http://localhost:*"), ","),
	})
}
