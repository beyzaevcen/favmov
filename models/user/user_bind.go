package user

import (
	"errors"
	db "fav-mov/db/sqlc"
	"net/http"
)

type UserBind struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Image    string `json:"image"`
	Password string `json:"password_hash"`
}

func (ru *UserBind) UserBindParams() db.RegisterUserParams {
	return db.RegisterUserParams{
		Name:         ru.Name,
		Image:        ru.Image,
		PasswordHash: ru.Password,
		Email:        ru.Email,
	}
}

func (ru *UserBind) Bind(r *http.Request) error {
	if len(ru.Name) < 3 {
		return errors.New("name must be at least 3 characters")
	}

	return nil
}
