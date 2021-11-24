package repositories

import (
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/database"
	"github.com/startup-of-zero-reais/COD-users-api/domain/entities"
	"gorm.io/gorm"
)

type (
	User struct {
		db *database.Database
	}
)

func NewUser(db *database.Database) *User {
	return &User{
		db: db,
	}
}

func (u *User) Get(ids []string, limit uint, offset uint) ([]entities.User, int) {
	var users []entities.User

	u.db.Conn.Limit(int(limit)).Offset(int(offset)).Find(&users, ids)

	return users, u.TotalRecords()
}

func (u *User) Search(search map[string]interface{}) ([]entities.User, int) {
	var users []entities.User

	u.db.Conn.Where(search).Find(&users)

	return users, u.TotalRecords()
}

func (u *User) Save(user *entities.User) *entities.User {
	if user.ID != "" {
		u.db.Conn.Where("user_id = ?", user.ID).Updates(user)
	} else {
		u.db.Conn.Create(user)
	}

	return user
}

func (u *User) Delete(id string) bool {
	r := u.db.Conn.Where("user_id = ?", id).Delete(new(entities.User))

	return r.RowsAffected > 0
}

func (u *User) TotalRecords() int {
	var find []entities.User
	var total int

	u.db.Conn.Select("user_id").FindInBatches(&find, 10000, func(tx *gorm.DB, batch int) error {
		total = int(tx.RowsAffected)
		return nil
	})

	return total
}
