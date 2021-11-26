package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type (
	// RecoverToken é a estrutura do token de recuperação de conta
	RecoverToken struct {
		ID        string    `json:"id" gorm:"column:recover_token_id;primaryKey;type:varchar(36);"`
		Token     string    `json:"token,omitempty"`
		Email     string    `json:"email,omitempty"`
		CreatedAt time.Time `json:"created_at"`
	}
)

// BeforeCreate é responsável por gerar um UUID para os tokens ao criá-los
func (r *RecoverToken) BeforeCreate(_ *gorm.DB) error {
	r.ID = uuid.New().String()

	return nil
}
