package repositories

import (
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/database"
	"github.com/startup-of-zero-reais/COD-users-api/domain/entities"
	"time"
)

type (
	// Token é a estrutura de repositório para anexar os métodos necessários
	Token struct {
		db *database.Database
	}
)

// NewToken instancia um Token
func NewToken(db *database.Database) *Token {
	return &Token{
		db: db,
	}
}

// Get recupera um entities.RecoverToken case exista.
// Pode retornar nil caso não encontre o valor na base de dados
func (t *Token) Get(id string) *entities.RecoverToken {
	token := &entities.RecoverToken{}

	t.db.Conn.Where("recover_token_id = ?", id).First(token)
	if token.ID == "" {
		return nil
	}

	return token
}

// Generate é o método que salva na base de dados um entities.RecoverToken
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
