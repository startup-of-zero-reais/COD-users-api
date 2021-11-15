package routes

import (
	"github.com/labstack/echo/v4"
	validatorsAdapter "github.com/startup-of-zero-reais/COD-users-api/adapters/http/validators"
	"github.com/startup-of-zero-reais/COD-users-api/domain/entities"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/services"
	"github.com/startup-of-zero-reais/COD-users-api/domain/ports/validators"
	"log"
	"net/http"
)

type (
	User struct {
		Routes    []*Route
		Service   services.UserService
		Validator validators.UserValidator
		Group     *echo.Group
	}
)

func NewUser(g *echo.Group) *User {
	return &User{
		Service:   nil,
		Group:     g,
		Validator: validatorsAdapter.NewUser(),
	}
}

func (u *User) List() {
	route := NewRoute(u.Group)

	route.register(func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "List Users ok!"})
	})

	u.register(route)
}

func (u *User) Create() {
	route := NewRoute(u.Group)
	route.Method = "POST"

	route.register(func(c echo.Context) error {
		err := u.validate(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, validationError("erro de validação", err))
		}

		return c.JSON(http.StatusCreated, map[string]string{"message": "fake created"})
	})

	u.register(route)
}

func (u *User) Register() {
	u.List()
	u.Create()

	for _, route := range u.Routes {
		route.registerMethod()
	}

	log.Println("done! user routes registered")
}

func (u *User) register(route *Route) {
	u.Routes = append(u.Routes, route)
}

func (u *User) validate(c echo.Context) []validators.Error {
	user := new(entities.User)
	user.New()

	err := c.Bind(user)
	if err != nil {
		return []validators.Error{
			{
				Field:   "Bind",
				Message: err.Error(),
			},
		}
	}

	errs := u.Validator.Validate(user)

	if errs != nil && len(errs) > 0 {
		return errs
	}

	return nil
}
