package repository

import (
	"crypto/sha512"
	"database/sql"
	"fmt"

	apiError "phrasebook-api/src/error"
	"phrasebook-api/src/model"
)

type UserRepository interface {
	CreateUser(*model.User) error
	GetUserByEmail(string) (*model.User, error)
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) CreateUser(user *model.User) error {
	existingUser, searchErr := u.GetUserByEmail(user.Email)

	if searchErr != nil {
		return searchErr
	}

	if existingUser != nil {
		return apiError.NewValidationError("email", "such email already exists")
	}

	user.Password = passwordHash(user.Password)

	result, insertErr := u.db.Exec(
		"INSERT INTO users (name, password, email) VALUES(?, ?, ?)",
		user.Name,
		user.Password,
		user.Email,
	)

	if insertErr != nil {
		return insertErr
	}

	userID, resultErr := result.LastInsertId()
	if resultErr != nil {
		return resultErr
	}

	user.ID = userID

	return nil
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
