package user

import (
	"net/http"

	"phrasebook-api/src/repository"

	"github.com/labstack/echo"
	"phrasebook-api/src/errors"
	"phrasebook-api/src/response"
)

func NewRegistrationHandler(user repository.UserRepository) *registrationHandler {
	return &registrationHandler{
		userRepository: user,
	}
}

type registrationHandler struct {
	userRepository repository.UserRepository
}

func (u *registrationHandler) InitRouting(router *echo.Echo) {
	router.POST("/v1/user/registration", u.registration)
}

func (u *registrationHandler) registration(c echo.Context) error {
	data := &RegistrationRequest{}
	if err := c.Bind(data); err != nil {
		return err
	}

	if err := data.Validation(); err != nil {
		return err
	}

	if user, _ := u.userRepository.GetUserByEmail(data.Email); user != nil {
		return errors.NewValidationError().AddFieldError("email", "user already exists")
	}

	if _, err := u.userRepository.CreateUser(data.Name, data.Password, data.Email); err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, response.Success())
	}
}
