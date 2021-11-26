package services

import (
	"errors"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/database"
	repositoriesAdapter "github.com/startup-of-zero-reais/COD-users-api/adapters/http/repositories"
	"github.com/startup-of-zero-reais/COD-users-api/domain/entities"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/repositories"
)

type (
	// User é a estrutura de serviço de usuários
	User struct {
		repo repositories.UserRepository
	}
)

// NewUser é o construtor de User
func NewUser(db *database.Database) *User {
	return &User{
		repo: repositoriesAdapter.NewUser(db),
	}
}

// O paginate é responsável por retornar a partir de qual registro será feita a busca
// no repositório de usuário
func (us *User) paginate(page uint, perPage uint) uint {
	offset := perPage
	if page >= 1 {
		offset = (page * perPage) - perPage
	}

	return offset
}

// List é o método responsável por listar os usuários e o total de usuários na base de dados
func (us *User) List(ids []string, page uint, perPage uint) ([]entities.User, int) {
	offset := us.paginate(page, perPage)
	users, total := us.repo.Get(ids, perPage, offset)
	var usersHiddenFields []entities.User

	for _, u := range users {
		(&u).HideSensitiveFields()
		usersHiddenFields = append(usersHiddenFields, u)
	}

	return usersHiddenFields, total
}

// Get recupera um único usuário baseado no seu ID
func (us *User) Get(id string) *entities.User {
	findId := []string{id}
	users, _ := us.repo.Get(findId, 1, 0)

	return &users[0]
}

// Create registra um novo usuário no repositório de usuário
func (us *User) Create(user *entities.User) (*entities.User, error) {
	isSetUser, _ := us.repo.Search(map[string]interface{}{"email": user.Email})
	if len(isSetUser) > 0 {
		return nil, errors.New("usuário já cadastrado")
	}

	user.NewPassword = user.Password
	err := user.HashPassword()
	if err != nil {
		return nil, err
	}

	createdUser := us.repo.Save(user)
	createdUser.HideSensitiveFields()

	return createdUser, nil
}

// Update atualiza o registro de um usuário no repositório
func (us *User) Update(id string, user *entities.User) (*entities.User, error) {
	currentUserResponse, _ := us.repo.Get([]string{id}, 1, 0)

	if len(currentUserResponse) == 0 {
		return nil, errors.New("usuário não encontrado")
	}
	currentUser := currentUserResponse[0]

	if user.Password != "" {
		if !currentUser.IsValidPassword(user.Password) {
			return nil, errors.New("credenciais inválidas")
		}

		err := user.HashPassword()

		if err != nil {
			return nil, err
		}
	} else {
		user.Password = currentUser.Password
	}

	user.ID = currentUser.ID
	user.CreatedAt = currentUser.CreatedAt

	if user.Email == "" {
		user.Email = currentUser.Email
	}
	if user.Type == ("") {
		user.Type = currentUser.Type
	}

	updatedUser := us.repo.Save(user)
	updatedUser.HideSensitiveFields()

	return updatedUser, nil
}

// Delete apaga o registro de um usuário no repositório
func (us *User) Delete(id string) bool {
	return us.repo.Delete(id)
}
