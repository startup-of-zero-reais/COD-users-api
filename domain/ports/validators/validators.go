package validators

import "github.com/startup-of-zero-reais/COD-users-api/domain/entities"

type (
	Error struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}

	UserValidator interface {
		Validate(user *entities.User) []Error
	}
)

func NewValidatorErrors(message error, fields ...string) []Error {
	return []Error{
		{
			Field:   fields[0],
			Message: message.Error(),
		},
	}
}
