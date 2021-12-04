package bearer_jwt

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/server/controllers/router"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/services"
	"github.com/startup-of-zero-reais/COD-users-api/domain/entities"
	"github.com/startup-of-zero-reais/COD-users-api/domain/utilities"
	"time"
)

// JwtHeaderConfig retorna o middleware que valida o jwt de autenticação
func JwtHeaderConfig() router.MiddlewareHandler {
	config := middleware.JWTConfig{
		Claims:     &services.JwtCustomClaims{},
		SigningKey: []byte(utilities.GetEnv("APP_SECRET", "tokn-secre7")),
	}

	return (router.MiddlewareHandler)(middleware.JWTWithConfig(config))
}

func JwtDecode(tokenString string) (*entities.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(utilities.GetEnv("APP_SECRET", "tokn-secre7")), nil
	})

	if err != nil {
		return nil, err
	}

	var user entities.User

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user.ID = claims["sub"].(string)
		user.Email = claims["email"].(string)
		user.Name = claims["name"].(string)
		user.Lastname = claims["lastname"].(string)
		user.Type = entities.UserType(claims["user_type"].(string))
		user.CreatedAt, _ = time.Parse(time.RFC3339, claims["created_at"].(string))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, claims["updated_at"].(string))
		user.HideSensitiveFields()

		return &user, nil
	}

	return nil, errors.New("invalid claim user")
}
