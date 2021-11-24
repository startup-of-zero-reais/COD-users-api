package services

import "github.com/startup-of-zero-reais/COD-users-api/domain/entities"

type (
	UserService interface {
		List(ids []string, page uint, perPage uint) ([]entities.User, int)
		Get(id string) *entities.User
		Create(user *entities.User) (*entities.User, error)
		Update(id string, user *entities.User) (*entities.User, error)
		Delete(id string) bool
	}

	AuthService interface {
		Get(email string) *entities.User
	}

	RecoverAccountService interface {
		SendEmail(email string) bool
		UpdatePassword(token, email, newPassword string) error
	}
)
