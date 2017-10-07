package repository

import (
	"crypto/sha512"
	"database/sql"
	"fmt"
	"phrasebook-api/src/model"
)

type UserRepository interface {
	CreateUser(name, password, email string) (*model.User, error)
	GetUserByEmail(string) (*model.User, error)
}

var _ UserRepository = (*userRepository)(nil)

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) CreateUser(name, password, email string) (*model.User, error) {
	password = passwordHash(password)

	result, err := u.db.Exec(
		"INSERT INTO users (name, password, email) VALUES(?, ?, ?)",
		name, password, email,
	)

	if err != nil {
		return nil, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	userModel := &model.User{
		ID:       userID,
		Name:     name,
		Password: password,
		Email:    email,
	}

	return userModel, nil
}

func (u *userRepository) GetUserByEmail(email string) (*model.User, error) {
	user := model.User{}
	row := u.db.QueryRow("SELECT id_user, name, email, password FROM users WHERE email=?", email)
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func passwordHash(password string) string {
	sha := sha512.New()
	sha.Write([]byte(password))
	return fmt.Sprintf("%x", sha.Sum(nil))
}
