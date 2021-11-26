package resources

import "github.com/startup-of-zero-reais/COD-users-api/domain/entities"

type (
	// Resource é a 'interface' necessária paga que seja criada a referência de Href
	// na entidade do domínio
	Resource interface {
		GetEmbedded()
	}

	// UserCollection é a 'interface' necessária para recuperar uma coleção de usuários
	// a partir de um slice de entities.User
	UserCollection interface {
		GetCollection([]entities.User) []interface{ Resource }
	}
)
