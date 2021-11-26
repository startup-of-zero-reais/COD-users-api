package validators

import "github.com/startup-of-zero-reais/COD-users-api/domain/entities"

type (
	// Error estrutura relacionada com o erro de validadores
	Error struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}

	// UserValidator 'interface' de validação de usuários
	UserValidator interface {
		Validate(user *entities.User) []Error
	}
)

// NewValidatorErrors é responsável por montar um slice de Error
func NewValidatorErrors(message error, fields ...string) []Error {
	return []Error{
		{
			Field:   fields[0],
			Message: message.Error(),
		},
	}
}
