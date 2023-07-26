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

	comment, err := store.AddComment(ctx, commentparams.AddCommentParams(userId))
	if err != nil {
		render.Render(w, r, status.ErrInternal(err))
		return
	}

	userInfo, err := store.GetImageAndNameOfUser(ctx, comment.UserID)
	if err != nil {
		render.Render(w, r, status.ErrNotFoundOrInternal(err))
		return
	}

	commentRow := db.GetCommentsRow{
		ID:        comment.ID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		Name:      userInfo.Name,
		Image:     userInfo.Image,
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, &commentRow)
}
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	store := ctx.Value(StoreKey).(*db.Store)
	userID := ctx.Value(IDKey).(int64)

	ID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	id, err := store.DeleteComment(ctx, db.DeleteCommentParams{
		ID:     ID,
		UserID: userID,
	})
	if err != nil {
		render.Render(w, r, status.ErrNotFoundOrInternal(err))
		return
	}

	render.Render(w, r, status.SuccessID(id, "Successfully deleted :)"))
}

func GetMyComments(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	store := ctx.Value(StoreKey).(*db.Store)
	userID := ctx.Value(IDKey).(int64)

	movieID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		render.Render(w, r, status.ErrBadRequest(err))
		return
	}

	comments, err := store.GetMyComments(ctx, db.GetMyCommentsParams{MovieID: movieID, UserID: userID})
	if err != nil {
		render.Render(w, r, status.ErrInternal(err))
		return
	}

	commentsPayload := utils.Map(comments, func(el db.GetMyCommentsRow) render.Renderer { return &el })

	render.Status(r, http.StatusOK)
	render.RenderList(w, r, utils.NewRenderList(commentsPayload))
}
