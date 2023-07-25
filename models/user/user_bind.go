package user

import (
	"errors"
	db "fav-mov/db/sqlc"
	"net/http"
)

type UserBind struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (ru *UserBind) UserBindParams() db.RegisterUserParams {
	return db.RegisterUserParams{
		Name:  ru.Name,
		Email: ru.Email,
	}
}

func (ru *UserBind) Bind(r *http.Request) error {
	if len(ru.Name) < 3 {
		return errors.New("name must be at least 3 characters")
	}

	return nil
}
