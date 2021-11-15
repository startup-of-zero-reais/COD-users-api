package entities

import "time"

type (
	UserType string

	User struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		Lastname  string    `json:"lastname"`
		Email     string    `json:"email"`
		Type      UserType  `json:"user_type"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

const (
	student = UserType("student")
	teacher = UserType("teacher")
)
