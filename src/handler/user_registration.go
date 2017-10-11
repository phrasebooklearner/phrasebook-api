package handler

import (
	"net/http"

	"phrasebook-api/src/repository"
	"phrasebook-api/src/validation"

	"github.com/labstack/echo"
)

type RegistrationRequest struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"-" form:"password"`
	Email    string `json:"email" form:"email"`
}

func NewRegistrationHandler(user repository.UserRepository) *registrationHandler {
	return &registrationHandler{
		userRepository: user,
	}
}

type registrationHandler struct{
	userRepository repository.UserRepository
}

func (u *registrationHandler) InitRouting(router *echo.Echo) {
	router.POST("/registration", u.registration)
}

func (u *registrationHandler) registration(c echo.Context) error {
	user := &RegistrationRequest{}
	if err := c.Bind(user); err != nil {
		return err
	}

	if err := validation.ValidateNewUser(user); err != nil {
		return err
	}

	//if err := u.userRepository.CreateUser(user); err != nil {
	//	return err
	//}

	return c.JSON(http.StatusOK, user)
}

//func ValidateNewUser(user *model.User) error {
//	if !govalidator.IsEmail(user.Email) {
//		return apiError.NewValidationError("email", "invalid email")
//	}
//
//	if !govalidator.StringLength(user.Name, "2", "255") {
//		return apiError.NewValidationError("name", "name is too short")
//	}
//
//	if !govalidator.StringLength(user.Password, "5", "255") {
//		return apiError.NewValidationError("password", "password is too short")
//	}
//
//	return nil
//}
