package model

type User struct {
	ID       int64  `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Password string `json:"-" form:"password"`
	Email    string `json:"email" form:"email"`
}
