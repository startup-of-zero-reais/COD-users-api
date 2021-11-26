package entities

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/startup-of-zero-reais/COD-users-api/domain/utilities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type (
	// UserType é apenas uma definição de tipo para User.Type
	UserType string

	// User é o formato da estrutura de usuário
	User struct {
		ID          string    `json:"id" gorm:"primaryKey;column:user_id;type:varchar(36);"`
		Name        string    `json:"name" gorm:"name" validate:"required"`
		Lastname    string    `json:"lastname" gorm:"lastname" validate:"required"`
		Email       string    `json:"email" gorm:"email;unique" validate:"required_if=ID ''|email"`
		Type        UserType  `json:"user_type" gorm:"user_type"`
		Password    string    `json:"password,omitempty" gorm:"password" validate:"required_with=Email|min=6"`
		NewPassword string    `json:"new_password,omitempty" gorm:"-" validate:"required_with=Password"`
		Href        string    `json:"_href,omitempty" gorm:"-" validate:"-"`
		CreatedAt   time.Time `json:"created_at" gorm:"created_at"`
		UpdatedAt   time.Time `json:"updated_at" gorm:"updated_at"`
	}
)

// BeforeCreate gera o UUID padrão antes de criar um usuário e atribui o tipo "student" se não houver definição
func (u *User) BeforeCreate(_ *gorm.DB) error {
	u.ID = uuid.New().String()

	if u.Type == "" {
		u.Type = student
	}

	return nil
}

// BeforeSave Atribui o tipo "student" se não houver definição
func (u *User) BeforeSave(_ *gorm.DB) error {
	if u.Type == "" {
		u.Type = student
	}

	return nil
}

// HashPassword recebe cria um hash de senha a partir do attr NewPassword do usuário
func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.NewPassword), 14)
	u.Password = string(bytes)
	return err
}

// IsValidPassword compara uma senha com o hash da Password do usuário
func (u *User) IsValidPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// HideSensitiveFields é responsável por simplesmente subscrever campos sensíveis do usuário por vazio para serem ocultados
func (u *User) HideSensitiveFields() {
	u.Password = ""
	u.NewPassword = ""
}

// GetEmbedded implementa um Resource para atribuir o Href do usuário com o link de direcionamento para o User
func (u *User) GetEmbedded() {
	if u.ID != "" {
		baseURL := utilities.GetEnv("APPLICATION_HOST", "http://localhost:8080")
		u.Href = fmt.Sprintf("%s/users?ids=%s&page=1&per_page=1", baseURL, u.ID)
		return
	}

	u.Href = ""
}

const (
	// student é uma constante para atribuir o tipo para o usuário
	student = UserType("student")
	// teacher é uma constante para atribuir o tipo para o usuário
	teacher = UserType("teacher")
)
