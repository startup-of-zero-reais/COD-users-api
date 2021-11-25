package services

import (
	"github.com/golang-jwt/jwt"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/database"
	repositoriesAdapter "github.com/startup-of-zero-reais/COD-users-api/adapters/http/repositories"
	"github.com/startup-of-zero-reais/COD-users-api/domain/entities"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/repositories"
)

type (
	Auth struct {
		repo repositories.UserRepository
	}

	JwtCustomClaims struct {
		Name  string `json:"name"`
		Email string `json:"email,omitempty"`
		jwt.StandardClaims
	}
)

func NewAuth(db *database.Database) *Auth {
	return &Auth{
		repo: repositoriesAdapter.NewUser(db),
	}
}

func (a *Auth) Get(email string) *entities.User {
	searchEmail := map[string]interface{}{
		"email": email,
	}
	users, _ := a.repo.Search(searchEmail)

	if len(users) <= 0 {
		return nil
	}

	return &users[0]
}
