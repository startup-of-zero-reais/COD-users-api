package services

import "github.com/startup-of-zero-reais/COD-users-api/domain/entities"

type (
	UserService interface {
		List(query map[string]interface{}, page uint, perPage uint) []*entities.User
		Get(id string) *entities.User
		Create(user *entities.User) *entities.User
		Update(id string, user *entities.User) *entities.User
		Delete(id string) *entities.User
	}
)
