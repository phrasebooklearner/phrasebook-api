// +build !unit

package repository

import (
	"testing"

	"phrasebook-api/src/config"
	"phrasebook-api/src/database"
	apiError "phrasebook-api/src/error"
	"phrasebook-api/src/model"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndSearchUser_Success(t *testing.T) {
	// arrange
	newUser := &model.User{
		Name: "petya",
		Password: "some_strange_pass",
		Email: "petya@vasya.com",
	}
	cfg := config.NewTestConfig()
	db := database.NewDBConnection(cfg.GetDatabaseDSN())
	rep := NewUserRepository(db)

	// act
	createErr := rep.CreateUser(newUser)
	searchUser, searchErr := rep.GetUserByEmail(newUser.Email)

	// assert
	assert.Nil(t, createErr)
	assert.Nil(t, searchErr)
	assert.Equal(t, *newUser, *searchUser)
}

func TestCreateUser_AlreadyExists(t *testing.T) {
	// arrange
	user1 := &model.User{
		Name: "petya",
		Password: "some_strange_pass",
		Email: "vasya@petya.com",
	}
	user2 := &model.User{
		Name: "vasya",
		Password: "some_strange_pass",
		Email: "vasya@petya.com",
	}
	cfg := config.NewTestConfig()
	db := database.NewDBConnection(cfg.GetDatabaseDSN())
	rep := NewUserRepository(db)

	// act
	first := rep.CreateUser(user1)
	second := rep.CreateUser(user2)
	err, ok := second.(apiError.ValidationError)

	// assert
	assert.Nil(t, first)
	assert.NotNil(t, second)
	assert.True(t, ok)
	assert.Equal(t, err.Field, "email")
	assert.Equal(t, err.Message, "such email already exists")
}