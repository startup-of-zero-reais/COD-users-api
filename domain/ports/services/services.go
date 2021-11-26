package services

import "github.com/startup-of-zero-reais/COD-users-api/domain/entities"

type (
	// UserService 'interface' para que os serviços de usuário funcionem corretamente
	UserService interface {
		List(ids []string, page uint, perPage uint) ([]entities.User, int)
		Get(id string) *entities.User
		Create(user *entities.User) (*entities.User, error)
		Update(id string, user *entities.User) (*entities.User, error)
		Delete(id string) bool
	}

	// AuthService 'interface' para o serviço de autenticação
	AuthService interface {
		Get(email string) *entities.User
	}

	// RecoverAccountService 'interface' para o serviço de recuperação de conta
	RecoverAccountService interface {
		SendEmail(email string) bool
		GetToken(id string) (*entities.RecoverToken, error)
		ValidateToken(token string) error
		UpdatePassword(email, newPassword string) error
	}
)
