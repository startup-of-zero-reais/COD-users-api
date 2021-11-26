package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/middlewares/x_api_key"
	"github.com/startup-of-zero-reais/COD-users-api/domain/utilities"
	"net/http"
)

// GenKey registra a rota para gerar chaves de api na aplicação
func GenKey() (string, func(ctx echo.Context) error, echo.MiddlewareFunc) {
	auth := x_api_key.NewXApiKey()

	return "/gen-api-keys", func(ctx echo.Context) error {
		application := ctx.Request().Header.Get("application")
		type secretBody struct {
			Secret string `json:"secret"`
		}

		secret := &secretBody{}

		if err := ctx.Bind(secret); err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		if sec := utilities.GetEnv("APP_SECRET", "secret0-123-0z"); secret.Secret != sec {
			return ctx.JSON(http.StatusForbidden, map[string]string{"message": "Forbidden"})
		}

		key, err := auth.GenerateApiKey(application)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		return ctx.JSON(http.StatusOK, map[string]string{"message": key, "application": application})
	}, (func(c echo.HandlerFunc) echo.HandlerFunc)(auth.CheckApplication())
}
