package router

import (
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/validators"
)

type (
	Errors struct {
		Error   *string `json:"error,omitempty"`
		Field   *string `json:"field,omitempty"`
		Message *string `json:"message,omitempty"`
	}

	HttpResponse struct {
		Message string             `json:"message"`
		Errors  []validators.Error `json:"errors,omitempty"`
	}
)

func HttpError(message error) HttpResponse {
	return HttpResponse{
		Message: message.Error(),
	}
}

func ValidationError(message string, err []validators.Error) HttpResponse {
	return HttpResponse{
		Message: message,
		Errors:  err,
	}
}
