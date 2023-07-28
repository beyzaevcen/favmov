package handler

import (
	db "fav-mov/db/sqlc"
	"fav-mov/models/status"
	"fav-mov/models/user"
	"fav-mov/utils"
	"net/http"

	"github.com/go-chi/render"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	userParams := &user.UserBind{}
	err := render.Bind(r, userParams)
	if err != nil {
		render.Render(w, r, status.ErrBadRequest(err))
		return
	}

	ctx := r.Context()
	store := ctx.Value(StoreKey).(*db.Store)

	usr, err := store.RegisterUser(ctx, userParams.UserBindParams())
	if err != nil {
		render.Render(w, r, status.ErrInternal(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, user.NewUserPayload(usr))
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	loginUser := &user.LoginUserBind{}
	err := render.Bind(r, loginUser)
	if err != nil {
		render.Render(w, r, status.ErrBadRequest(err))
		return
	}

	ctx := r.Context()
	store := ctx.Value(StoreKey).(*db.Store)

	userData, err := store.GetUserByEmail(ctx, loginUser.Email)
	if err != nil {
		render.Render(w, r, status.ErrInternal(err))
		return
	}

	if err != nil {
		render.Render(w, r, status.ErrInternal(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, user.NewUserPayloadFromUser(userData))
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	store := ctx.Value(StoreKey).(*db.Store)

	users, err := store.GetUsers(ctx)
	if err != nil {
		render.Render(w, r, status.ErrInternal(err))
		return
	}

	userPayload := utils.Map(users, func(el db.User) render.Renderer { return &el })

	render.Status(r, http.StatusOK)
	render.RenderList(w, r, utils.NewRenderList(userPayload))

}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	store := ctx.Value(StoreKey).(*db.Store)
	userID := ctx.Value(IDKey).(int64)

	userData, err := store.GetUserById(ctx, userID)
	if err != nil {
		render.Render(w, r, status.ErrNotFoundOrInternal(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, user.NewUserPayloadFromUser(userData))

}
