package services

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/database"
	repositoriesAdapter "github.com/startup-of-zero-reais/COD-users-api/adapters/http/repositories"
	"github.com/startup-of-zero-reais/COD-users-api/domain/entities"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/repositories"
	"github.com/startup-of-zero-reais/COD-users-api/domain/utilities"
	"log"
	"net/smtp"
	"time"
)

type (
	// RecoverAccount é a estrutura de serviço de recuperação de conta
	RecoverAccount struct {
		repo      repositories.UserRepository
		tokenRepo repositories.TokenRepository
	}
)

// NewRecoverAccount é o construtor de RecoverAccount
func NewRecoverAccount(db *database.Database) *RecoverAccount {
	return &RecoverAccount{
		repo:      repositoriesAdapter.NewUser(db),
		tokenRepo: repositoriesAdapter.NewToken(db),
	}
}

// O método genToken gera um token novo baseado no e-mail do usuário
func (r *RecoverAccount) genToken(email string) (*entities.RecoverToken, error) {
	claims := JwtCustomClaims{
		Name:  "recover-token",
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 4).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSecret := utilities.GetEnv("RECOVER_SECRET", "123456")
	t, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return nil, err
	}

	generatedToken := r.tokenRepo.Generate(t, email)
	if generatedToken == nil {
		return nil, errors.New("falha ao gerar token")
	}

	return generatedToken, nil
}

// SendEmail envia um e-mail de recuperação de conta para o usuário que corresponde
// ao e-mail do parâmetro, caso exista tal usuário
func (r *RecoverAccount) SendEmail(email string) bool {
	searchEmail := map[string]interface{}{
		"email": email,
	}
	users, _ := r.repo.Search(searchEmail)

	if len(users) <= 0 {
		log.Println("credenciais inválidas")
		return false
	}

	token, err := r.genToken(email)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	textMessage := fmt.Sprintf(
		"Você solicitou a recuperação de conta.\n\nCaso não tenha sido você apenas desconsidere.\n"+
			"Segue o link de recuperação: %s/redefinir-senha?token=%s\n\n"+
			"OBS.: Este link vale por 4 horas.",
		utilities.GetEnv("RECOVER_BASE_URL", "http://localhost:3000"),
		token.ID,
	)

	err = sendMail(email, "Recuperação de conta | Code Craft", textMessage)
	if err != nil {
		return false
	}

	return true
}

// GetToken recupera um entities.RecoverToken da base caso o ID esteja correto
// e retorna um erro quando não existe o token
func (r *RecoverAccount) GetToken(id string) (*entities.RecoverToken, error) {
	token := r.tokenRepo.Get(id)

	if token == nil {
		return nil, errors.New("token não encontrado ou inválido")
	}

	return token, nil
}

// ValidateToken decodifica o token e checa se é um token válido
func (r *RecoverAccount) ValidateToken(token string) error {
	parsedToken, err := jwt.ParseWithClaims(token, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		tokenSecret := utilities.GetEnv("RECOVER_SECRET", "123456")
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return err
	}

	if claims, ok := parsedToken.Claims.(*JwtCustomClaims); ok && parsedToken.Valid {
		fmt.Printf("%v %v", claims.Email, claims.StandardClaims.ExpiresAt)
	}

	return nil
}

// UpdatePassword atualiza a senha do usuário que corresponde ao e-mail
func (r *RecoverAccount) UpdatePassword(email, newPassword string) error {
	searchField := map[string]interface{}{
		"email": email,
	}
	users, _ := r.repo.Search(searchField)
	if len(users) <= 0 {
		return errors.New("nenhum usuário corresponde ao e-mail")
	}

	user := &users[0]
	user.NewPassword = newPassword
	err := user.HashPassword()
	if err != nil {
		return err
	}

	r.repo.Save(user)

	return nil
}

// O sendMail é um helper para enviar o e-mail com o formato correto
func sendMail(toMail, subject, theMessage string) error {
	from := utilities.GetEnv("SMTP_EMAIL", "jean.carlosmolossi@gmail.com")
	password := utilities.GetEnv("SMTP_PASSWORD", "123456")
	to := []string{toMail}

	smtpHost := utilities.GetEnv("SMTP_HOST", "smtp.gmail.com")
	smtpPort := utilities.GetEnv("SMTP_PORT", "587")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	textMessage := fmt.Sprintf(
		"To: %s\r\nSubject: %s\r\n\r\n%s",
		to,
		subject,
		theMessage,
	)

	message := []byte(textMessage)

	addr := fmt.Sprintf("%s:%s", smtpHost, smtpPort)
	err := smtp.SendMail(addr, auth, from, to, message)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
