package validators

import (
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ptBRTranslations "github.com/go-playground/validator/v10/translations/pt_BR"
	"github.com/startup-of-zero-reais/COD-users-api/domain/entities"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/validators"
	"log"
	"strings"
)

type (
	User struct {
		validate *validator.Validate
		trans    ut.Translator
	}
)

func NewUser() *User {
	ptBR := pt_BR.New()
	uni := ut.New(ptBR, ptBR)

	trans, _ := uni.GetTranslator("pt_BR")

	validate := validator.New()
	err := ptBRTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		log.Fatalln("Erro ao gerar traduções")
		return nil
	}

	return &User{
		validate: validate,
		trans:    trans,
	}
}

func (u *User) Validate(user *entities.User) []validators.Error {
	err := u.validate.Struct(user)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return []validators.Error{
				{
					Field:   "",
					Message: err.Error(),
				},
			}
		}

		var errors []validators.Error
		errs := err.(validator.ValidationErrors)
		translatedErrs := errs.Translate(u.trans)

		counter := 0
		for key, err := range translatedErrs {
			counter++

			structs := strings.Split(key, ".")
			field := structs[len(structs)-1]

			_err := validators.Error{
				Field:   field,
				Message: err,
			}

			errors = append(errors, _err)
		}

		return errors
	}

	return nil
}