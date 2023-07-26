package comment

import (
	db "fav-mov/db/sqlc"

	"net/http"
)

type CommentBind struct {
	Content string `json:"content"`
}

func (cb *CommentBind) AddCommentParams(userId int64, movieId int64) db.AddCommentParams {
	return db.AddCommentParams{
		UserID:  userId,
		MovieID: movieId,
		Content: cb.Content,
	}
}

func (cb *CommentBind) Bind(r *http.Request) error {
	return nil
}
