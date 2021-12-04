package services

import (
	"github.com/golang-jwt/jwt"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/database"
	repositoriesAdapter "github.com/startup-of-zero-reais/COD-users-api/adapters/http/repositories"
	"github.com/startup-of-zero-reais/COD-users-api/domain/entities"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/repositories"
	"time"
)

type (
	// Auth é a estrutura do serviço de autenticação
	Auth struct {
		repo repositories.UserRepository
	}

	// JwtCustomClaims é a estrutura de recuperação de tokens JWT personalizada
	JwtCustomClaims struct {
		Name      string            `json:"name"`
		Lastname  string            `json:"lastname"`
		Email     string            `json:"email,omitempty"`
		Type      entities.UserType `json:"user_type"`
		CreatedAt time.Time         `json:"created_at"`
		UpdatedAt time.Time         `json:"updated_at"`
		jwt.StandardClaims
	}
)

// NewAuth é o construtor de Auth
func NewAuth(db *database.Database) *Auth {
	return &Auth{
		repo: repositoriesAdapter.NewUser(db),
	}
}

// Get é o método para recuperar um usuário a partir de seu email
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
