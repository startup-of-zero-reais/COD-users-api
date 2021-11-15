package routes

import (
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/validators"
)

type (
	Errors struct {
		Error   *string `json:"error,omitempty"`
		Field   *string `json:"field,omitempty"`
		Message *string `json:"message,omitempty"`
	}

	httpResponse struct {
		Message string             `json:"message"`
		Errors  []validators.Error `json:"errors,omitempty"`
	}
)

func httpError(message string) httpResponse {
	return httpResponse{
		Message: message,
	}
}

func validationError(message string, err []validators.Error) httpResponse {
	return httpResponse{
		Message: message,
		Errors:  err,
	}
}
