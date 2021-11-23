package resources

import "github.com/startup-of-zero-reais/COD-users-api/domain/entities"

type (
	Resource interface {
		GetEmbedded()
	}

	UserCollection interface {
		GetCollection([]entities.User) []interface{ Resource }
	}
)
