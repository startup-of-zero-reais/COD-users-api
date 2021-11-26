package router

import (
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/validators"
)

type (
	// Errors é a estrutura de erros de validação montadas no roteamento da API
	Errors struct {
		Error   *string `json:"error,omitempty"`
		Field   *string `json:"field,omitempty"`
		Message *string `json:"message,omitempty"`
	}

	// HttpResponse é o modelo de resposta Http da API
	HttpResponse struct {
		Message string             `json:"message"`
		Errors  []validators.Error `json:"errors,omitempty"`
	}
)

// HttpError é um proxy para resultar uma mensagem de erro
func HttpError(message error) HttpResponse {
	return HttpResponse{
		Message: message.Error(),
	}
}

// ValidationError é um proxy para resultar uma mensagem de erro quando o erro é de validação
func ValidationError(message string, err []validators.Error) HttpResponse {
	return HttpResponse{
		Message: message,
		Errors:  err,
	}
}
