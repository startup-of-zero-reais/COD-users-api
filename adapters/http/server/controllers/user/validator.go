package user

import (
	"github.com/labstack/echo/v4"
	"github.com/startup-of-zero-reais/COD-users-api/domain/entities"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/validators"
)

func (u *User) validate(c echo.Context) (*entities.User, []validators.Error) {
	user := new(entities.User)

	err := c.Bind(user)
	if err != nil {
		return nil, validators.NewValidatorErrors(err, "Bind")
	}

	errs := u.Validator.Validate(user)

	if errs != nil && len(errs) > 0 {
		return nil, errs
	}

	return user, nil
}
