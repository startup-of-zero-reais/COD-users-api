package repositories

import "github.com/startup-of-zero-reais/COD-users-api/domain/entities"

type (
	UserRepository interface {
		Get(ids []string, limit uint, offset uint) ([]entities.User, int)
		Search(search map[string]interface{}) ([]entities.User, int)
		Save(user *entities.User) *entities.User
		Delete(id string) bool
	}
)
