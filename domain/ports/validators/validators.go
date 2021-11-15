package validators

import "github.com/startup-of-zero-reais/COD-users-api/domain/entities"

type (
	UserValidator interface {
		Validate(user *entities.User) error
	}
)
