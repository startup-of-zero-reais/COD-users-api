package repositories

import (
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/database"
	"github.com/startup-of-zero-reais/COD-users-api/domain/entities"
	"gorm.io/gorm"
)

type (
	// User é a estrutura de repositório para anexar os métodos
	User struct {
		db *database.Database
	}
)

// NewUser é o construtor para User
func NewUser(db *database.Database) *User {
	return &User{
		db: db,
	}
}

// Get é o método responsável por buscar um ou mais entities.User na base de dados
// e retornar um slice de usuários e o total de registros na base de dados
func (u *User) Get(ids []string, limit uint, offset uint) ([]entities.User, int) {
	var users []entities.User

	u.db.Conn.Limit(int(limit)).Offset(int(offset)).Find(&users, ids)

	return users, u.TotalRecords()
}

// Search é um método para procurar por registros na base de dados que correspondam a busca
//
// Pode retornar um slice vazio e o total de registros na base de dados
func (u *User) Search(search map[string]interface{}) ([]entities.User, int) {
	var users []entities.User

	u.db.Conn.Where(search).Find(&users)

	return users, u.TotalRecords()
}

// Save é o método que salva alterações ou novos registros na base de dados
func (u *User) Save(user *entities.User) *entities.User {
	if user.ID != "" {
		u.db.Conn.Where("user_id = ?", user.ID).Updates(user)
	} else {
		u.db.Conn.Create(user)
	}

	return user
}

// Delete serve para deletar registros a partir do id
func (u *User) Delete(id string) bool {
	r := u.db.Conn.Where("user_id = ?", id).Delete(new(entities.User))

	return r.RowsAffected > 0
}

// TotalRecords é responsável por buscar na base de dados quantos registros existem.
// Essa busca é executada em batches para melhorar a performance
func (u *User) TotalRecords() int {
	var find []entities.User
	var total int

	u.db.Conn.Select("user_id").FindInBatches(&find, 10000, func(tx *gorm.DB, batch int) error {
		total = int(tx.RowsAffected)
		return nil
	})

	return total
}
