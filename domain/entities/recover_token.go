package entities

import "time"

type (
	RecoverToken struct {
		ID        string    `json:"id" gorm:"column:recover_token_id;primaryKey;type:varchar(36);"`
		Token     string    `json:"token,omitempty"`
		Email     string    `json:"email,omitempty"`
		CreatedAt time.Time `json:"created_at"`
	}
)
