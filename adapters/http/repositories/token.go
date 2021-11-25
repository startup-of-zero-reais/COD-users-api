package repositories

import (
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/database"
	"github.com/startup-of-zero-reais/COD-users-api/domain/entities"
	"time"
)

type (
	Token struct {
		db *database.Database
	}
)

func NewToken(db *database.Database) *Token {
	return &Token{
		db: db,
	}
}

func (t *Token) Get(id string) *entities.RecoverToken {
	token := &entities.RecoverToken{}

	t.db.Conn.Where("recover_token_id = ?", id).First(token)
	if token.ID == "" {
		return nil
	}

	return token
}

func (t *Token) Generate(token, email string) *entities.RecoverToken {
	recoverToken := &entities.RecoverToken{
		Token:     token,
		Email:     email,
		CreatedAt: time.Now(),
	}

	res := t.db.Conn.Create(recoverToken)
	if res.Error != nil {
		return nil
	}

	return recoverToken
}
