package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// HealthCheck é a rota de verificação de saúde da aplicação
func HealthCheck() (string, func(ctx echo.Context) error) {
	return "/healthcheck", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]string{"message": "health"})
	}
}
