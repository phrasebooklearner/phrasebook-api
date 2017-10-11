package validation

import (
	apiError "phrasebook-api/src/response"
	"phrasebook-api/src/model"

	"github.com/asaskevich/govalidator"
)

func ValidateNewUser(user *model.User) error {
	if !govalidator.IsEmail(user.Email) {
		return apiError.NewValidationError("email", "invalid email")
	}

	if !govalidator.StringLength(user.Name, "2", "255") {
		return apiError.NewValidationError("name", "name is too short")
	}

	if !govalidator.StringLength(user.Password, "5", "255") {
		return apiError.NewValidationError("password", "password is too short")
	}

	return nil
}
