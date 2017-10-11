package validation

import (
	"testing"

	apiError "phrasebook-api/src/response"
	"phrasebook-api/src/model"

	"github.com/stretchr/testify/assert"
)

func TestValidateNewUser_EmptyEmail(t *testing.T) {
	// arrange
	user := &model.User{
		Password: "validpassword",
		Name:     "Name",
		Email:    "",
	}
	// act
	err := ValidateNewUser(user)
	valErr, ok := err.(apiError.ValidationError)
	// assert
	assert.NotNil(t, err)
	assert.True(t, ok)
	assert.Equal(t, valErr.Field, "email")
}

func TestValidateNewUser_InvalidEmail(t *testing.T) {
	// arrange
	user := &model.User{
		Password: "validpassword",
		Name:     "Name",
		Email:    "invalidemail",
	}
	// act
	err := ValidateNewUser(user)
	valErr, ok := err.(apiError.ValidationError)
	// assert
	assert.NotNil(t, err)
	assert.True(t, ok)
	assert.Equal(t, valErr.Field, "email")
}

func TestValidateNewUser_ShortName(t *testing.T) {
	// arrange
	user := &model.User{
		Password: "validpassword",
		Name:     "n",
		Email:    "validemail@gmail.com",
	}
	// act
	err := ValidateNewUser(user)
	valErr, ok := err.(apiError.ValidationError)
	// assert
	assert.NotNil(t, err)
	assert.True(t, ok)
	assert.Equal(t, valErr.Field, "name")
}

func TestValidateNewUser_ShortPassword(t *testing.T) {
	// arrange
	user := &model.User{
		Password: "1234",
		Name:     "Name",
		Email:    "validemail@gmail.com",
	}
	// act
	err := ValidateNewUser(user)
	valErr, ok := err.(apiError.ValidationError)
	// assert
	assert.NotNil(t, err)
	assert.True(t, ok)
	assert.Equal(t, valErr.Field, "password")
}
