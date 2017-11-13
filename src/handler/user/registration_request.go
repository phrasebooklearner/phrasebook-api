package user

import (
	"github.com/asaskevich/govalidator"
	"phrasebook-api/src/errors"
)

type RegistrationRequest struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"-" form:"password"`
	Email    string `json:"email" form:"email"`
}

func (r RegistrationRequest) Validation() error {
	err := errors.NewValidationError()

	if !govalidator.IsEmail(r.Email) {
		err.AddFieldError("email", "invalid email")
	}

	if !govalidator.StringLength(r.Name, "2", "255") {
		err.AddFieldError("name", "name is too short")
	}

	if !govalidator.StringLength(r.Password, "5", "255") {
		err.AddFieldError("password", "password is too short")
	}

	if !err.Empty() {
		return err
	}

	return nil
}
