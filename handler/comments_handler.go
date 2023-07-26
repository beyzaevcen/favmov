package handler

import (
	db "fav-mov/db/sqlc"
	"fav-mov/models/comment"
	"fav-mov/models/status"
	"fav-mov/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func GetComments(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	store := ctx.Value(StoreKey).(*db.Store)

	movieID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		render.Render(w, r, status.ErrBadRequest(err))
		return
	}

	comments, err := store.GetComments(ctx, movieID)
	if err != nil {
		render.Render(w, r, status.ErrInternal(err))
		return
	}

	commentsPayload := utils.Map(comments, func(el db.GetCommentsRow) render.Renderer { return &el })

	render.Status(r, http.StatusOK)
	render.RenderList(w, r, utils.NewRenderList(commentsPayload))
}
func AddComment(w http.ResponseWriter, r *http.Request) {
	commentparams := &comment.CommentBind{}
	err := render.Bind(r, commentparams)
	if err != nil {
		render.Render(w, r, status.ErrBadRequest(err))
		return
	}

	ctx := r.Context()
	store := ctx.Value(StoreKey).(*db.Store)
	userId := ctx.Value(IDKey).(int64) // user id yi tokenın içinden alıyoruz

	movieId, err := strconv.ParseInt(chi.URLParam(r, "movie_id"), 10, 64)
	if err != nil {
		render.Render(w, r, status.ErrBadRequest(err))
		return
	}

	comment, err := store.AddComment(ctx, commentparams.AddCommentParams(userId, movieId))
	if err != nil {
		render.Render(w, r, status.ErrInternal(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, &comment)
}
