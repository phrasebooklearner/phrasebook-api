// +build !unit

package repository

import (
	"testing"

	"phrasebook-api/src/config"
	"phrasebook-api/src/database"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndSearchUser_Success(t *testing.T) {
	// arrange
	name, password, email := getUserArguments("petya")
	cfg := config.NewTestConfig()
	db := database.NewDBConnection(cfg.GetDatabaseDSN())
	rep := NewUserRepository(db)

	// act
	newUser, createErr := rep.CreateUser(name, password, email)
	searchUser, searchErr := rep.GetUserByEmail(newUser.Email)

	// assert
	assert.Nil(t, createErr)
	assert.Nil(t, searchErr)
	assert.Equal(t, *newUser, *searchUser)
	assert.Equal(t, searchUser.Name, name)
	assert.Equal(t, searchUser.Password, passwordHash(password))
	assert.Equal(t, searchUser.Email, email)
}

func TestCreateUser_AlreadyExists(t *testing.T) {
	// arrange
	name, password, email := getUserArguments("vasya")
	cfg := config.NewTestConfig()
	db := database.NewDBConnection(cfg.GetDatabaseDSN())
	rep := NewUserRepository(db)

	// act
	user1, err1 := rep.CreateUser(name, password, email)
	user2, err2 := rep.CreateUser(name, password, email)

	// assert
	assert.NotNil(t, user1)
	assert.Nil(t, err1)
	assert.Nil(t, user2)
	assert.NotNil(t, err2)
}

func getUserArguments(name string) (string, string, string) {
	return name, "qwerty", name + "@gmail.com"
}