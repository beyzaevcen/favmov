package user

import (
	db "fav-mov/db/sqlc"
	"net/http"
	"time"
)

type UserPayload struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUserPayload(data db.RegisterUserRow, token string) *UserPayload {
	return &UserPayload{
		ID:        data.ID,
		Name:      data.Name,
		Image:     data.Image,
		Email:     data.Email,
		Token:     token,
		CreatedAt: data.CreatedAt,
	}
}

func NewUserPayloadFromUser(data db.User, token string) *UserPayload {
	return &UserPayload{
		ID:        data.ID,
		Name:      data.Name,
		Image:     data.Image,
		Email:     data.Email,
		Token:     token,
		CreatedAt: data.CreatedAt,
	}
}

func (up *UserPayload) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
