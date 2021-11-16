package services

import (
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/database"
	repositoriesAdapter "github.com/startup-of-zero-reais/COD-users-api/adapters/http/repositories"
	"github.com/startup-of-zero-reais/COD-users-api/domain/entities"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/repositories"
	"log"
)

type (
	User struct {
		repo repositories.UserRepository
	}
)

func NewUser(db *database.Database) *User {
	return &User{
		repo: repositoriesAdapter.NewUser(db),
	}
}

func (us *User) paginate(page uint, perPage uint) uint {
	offset := perPage
	if page >= 1 {
		offset = (page * perPage) - perPage
	}

	return offset
}

func (us *User) List(ids []string, page uint, perPage uint) []entities.User {
	offset := us.paginate(page, perPage)
	users := us.repo.Get(ids, perPage, offset)

	return users
}

func (us *User) Get(id string) *entities.User {
	findId := []string{id}
	users := us.repo.Get(findId, 1, 0)

	return &users[0]
}

func (us *User) Create(user *entities.User) *entities.User {
	isSetUser := us.repo.Get([]string{user.ID}, 1, 0)
	if len(isSetUser) > 0 {
		log.Fatalln("usuário ja cadastrado")
		return nil
	}

	createdUser := us.repo.Save(user)

	return createdUser
}

func (us *User) Update(id string, user *entities.User) *entities.User {
	return nil
}

func (us *User) Delete(id string) *entities.User {
	return nil
}
