package repositories

import (
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/database"
	"github.com/startup-of-zero-reais/COD-users-api/domain/entities"
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

func (u *User) Get(ids []string, limit uint, offset uint) []entities.User {
	var users []entities.User

	u.db.Conn.Find(&users, ids).Limit(int(limit)).Offset(int(offset))

	return users
}

func (u *User) Search(search map[string]interface{}) []entities.User {
	var users []entities.User

	u.db.Conn.Where(search).Find(&users)

	return users
}

func (u *User) Save(user *entities.User) *entities.User {
	if user.ID != "" {
		u.db.Conn.Where("id = ?", user.ID).Updates(user)
	} else {
		u.db.Conn.Create(user)
	}

	return user
}

func (u *User) Delete(id string) *entities.User {
	return nil
}
