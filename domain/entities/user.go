package entities

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type (
	UserType string

	User struct {
		ID          string    `json:"id" gorm:"primaryKey;user_id;type:varchar(36);"`
		Name        string    `json:"name" gorm:"name" validate:"required"`
		Lastname    string    `json:"lastname" gorm:"lastname" validate:"required"`
		Email       string    `json:"email" gorm:"email;unique" validate:"required_if=ID ''|email"`
		Type        UserType  `json:"user_type" gorm:"user_type"`
		Password    string    `json:"password,omitempty" gorm:"password" validate:"required_with=Email|min=6"`
		NewPassword string    `json:"new_password,omitempty" gorm:"-" validate:"required_with=Password"`
		CreatedAt   time.Time `json:"created_at" gorm:"created_at"`
		UpdatedAt   time.Time `json:"updated_at" gorm:"updated_at"`
	}
)

func (u *User) BeforeCreate(_ *gorm.DB) error {
	u.ID = uuid.New().String()

	if u.Type == "" {
		u.Type = student
	}

	return nil
}

func (u *User) BeforeSave(_ *gorm.DB) error {
	if u.Type == "" {
		u.Type = student
	}

	return nil
}

func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.NewPassword), 14)
	u.Password = string(bytes)
	return err
}

func (u *User) IsValidPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) HideSensitiveFields() {
	u.Password = ""
	u.NewPassword = ""
}

const (
	student = UserType("student")
	_       = UserType("teacher")
)
