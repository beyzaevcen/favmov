package comment

import (
	db "fav-mov/db/sqlc"

	"net/http"
)

type CommentBind struct {
	MovieID int64  `json:"movie_id"`
	Content string `json:"content"`
}

func (cb *CommentBind) AddCommentParams(userId int64) db.AddCommentParams {
	return db.AddCommentParams{
		UserID:  userId,
		MovieID: cb.MovieID,
		Content: cb.Content,
	}
}

func (cb *CommentBind) Bind(r *http.Request) error {
	return nil
}
