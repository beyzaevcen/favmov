package user

import (
	"errors"
	db "fav-mov/db/sqlc"
	"net/http"
)

type UserBind struct {
	Name       string `json:"name"`
	FirebaseID string `json:"uid"`
	Email      string `json:"email"`
	Image      string `json:"image"`
}

func (ru *UserBind) UserBindParams() db.RegisterUserParams {
	return db.RegisterUserParams{
		Name:        ru.Name,
		FirebaseUid: ru.FirebaseID,
		Image:       ru.Image,
		Email:       ru.Email,
	}
}

func (ru *UserBind) Bind(r *http.Request) error {
	if len(ru.Name) < 3 {
		return errors.New("name must be at least 3 characters")
	}

	return nil
}
