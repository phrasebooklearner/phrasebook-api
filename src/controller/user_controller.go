package controller

import (
	"net/http"

	"phrasebook-api/src/model"
	"phrasebook-api/src/repository"
	"phrasebook-api/src/validation"

	"github.com/labstack/echo"
)

func NewUserController(user repository.UserRepository) *userController {
	return &userController{
		userRepository: user,
	}
}

type userController struct{
	userRepository repository.UserRepository
}

func (u *userController) InitRouting(router *echo.Echo) {
	router.POST("/registration", u.createUser)
}

func (u *userController) createUser(c echo.Context) error {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		return err
	}

	if err := validation.ValidateNewUser(user); err != nil {
		return err
	}

	if err := u.userRepository.CreateUser(user); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
