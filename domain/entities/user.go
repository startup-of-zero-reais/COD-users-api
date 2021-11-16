package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type (
	UserType string

	User struct {
		ID        string    `json:"id" gorm:"primaryKey;user_id;type:varchar(36);"`
		Name      string    `json:"name" gorm:"name" validate:"required"`
		Lastname  string    `json:"lastname" gorm:"lastname" validate:"required"`
		Email     string    `json:"email" gorm:"email;unique" validate:"required,email"`
		Type      UserType  `json:"user_type" gorm:"user_type"`
		CreatedAt time.Time `json:"created_at" gorm:"created_at;autoCreateTime"`
		UpdatedAt time.Time `json:"updated_at" gorm:"updated_at;autoUpdateTime"`
	}
)

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()

	if u.Type == "" {
		u.Type = student
	}

	return
}

const (
	student = UserType("student")
	teacher = UserType("teacher")
)
