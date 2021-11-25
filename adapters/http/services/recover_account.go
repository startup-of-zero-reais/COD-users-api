package services

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/database"
	repositoriesAdapter "github.com/startup-of-zero-reais/COD-users-api/adapters/http/repositories"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/repositories"
	"github.com/startup-of-zero-reais/COD-users-api/domain/utilities"
	"log"
	"net/smtp"
	"time"
)

type (
	RecoverAccount struct {
		repo      repositories.UserRepository
		tokenRepo repositories.TokenRepository
	}
)

func NewRecoverAccount(db *database.Database) *RecoverAccount {
	return &RecoverAccount{
		repo: repositoriesAdapter.NewUser(db),
	}
}

func (r *RecoverAccount) genToken(email string) error {
	claims := JwtCustomClaims{
		Name:  "recover-token",
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 4).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("rec0v3r-seCr37"))
	if err != nil {
		return err
	}

	r.tokenRepo.Generate(t)

	return nil
}

func (r *RecoverAccount) SendEmail(email string) bool {
	searchEmail := map[string]interface{}{
		"email": email,
	}
	users, _ := r.repo.Search(searchEmail)

	if len(users) <= 0 {
		log.Println("credenciais inválidas")
		return false
	}

	from := utilities.GetEnv("SMTP_EMAIL", "jean.carlosmolossi@gmail.com")
	password := utilities.GetEnv("SMTP_PASSWORD", "123456")
	to := []string{email}

	smtpHost := utilities.GetEnv("SMTP_HOST", "smtp.gmail.com")
	smtpPort := utilities.GetEnv("SMTP_PORT", "587")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	textMessage := fmt.Sprintf(
		"To: %s\r\nSubject: %s\r\n\r\n%s",
		to,
		"Assunto do e-mail",
		"Mensagem de fato que vai pro e-mail",
	)

	message := []byte(textMessage)

	addr := fmt.Sprintf("%s:%s", smtpHost, smtpPort)
	err := smtp.SendMail(addr, auth, from, to, message)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	return true

}

func (r *RecoverAccount) UpdatePassword(token, email, newPassword string) error {

	return nil
}
