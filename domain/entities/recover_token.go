package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type (
	RecoverToken struct {
		ID        string    `json:"id" gorm:"column:recover_token_id;primaryKey;type:varchar(36);"`
		Token     string    `json:"token,omitempty"`
		Email     string    `json:"email,omitempty"`
		CreatedAt time.Time `json:"created_at"`
	}
)

func (r *RecoverToken) BeforeCreate(_ *gorm.DB) error {
	r.ID = uuid.New().String()

	return nil
}
