package user

import "net/http"

type LoginUserBind struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ru *LoginUserBind) Bind(r *http.Request) error {
	return nil
}
