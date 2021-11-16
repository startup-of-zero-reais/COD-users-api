package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/startup-of-zero-reais/COD-users-api/adapters/http/database"
	servicesAdapter "github.com/startup-of-zero-reais/COD-users-api/adapters/http/services"
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

func NewUser(g *echo.Group, db *database.Database) *User {
	return &User{
		Service:   servicesAdapter.NewUser(db),
		Validator: validatorsAdapter.NewUser(),
		Group:     g,
	}
}

func (u *User) List() {
	route := NewRoute(u.Group)

	route.register(func(c echo.Context) error {
		users := u.Service.List([]string{}, 1, 10)
		return c.JSON(http.StatusOK, users)
	})

	u.register(route)
}

func (u *User) Create() {
	route := NewRoute(u.Group)
	route.Method = "POST"

	route.register(func(c echo.Context) error {
		user, validateErr := u.validate(c)
		if validateErr != nil {
			return c.JSON(http.StatusBadRequest, validationError("erro de validação", validateErr))
		}

		createdUser := u.Service.Create(user)

		return c.JSON(http.StatusCreated, createdUser)
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

func (u *User) validate(c echo.Context) (*entities.User, []validators.Error) {
	user := new(entities.User)
	user.New()

	err := c.Bind(user)
	if err != nil {
		return nil, []validators.Error{
			{
				Field:   "Bind",
				Message: err.Error(),
			},
		}
	}

	errs := u.Validator.Validate(user)

	if errs != nil && len(errs) > 0 {
		return nil, errs
	}

	return user, nil
}
