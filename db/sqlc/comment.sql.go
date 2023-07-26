// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: comment.sql

package db

import (
	"context"
	"time"
)

const addComment = `-- name: AddComment :one
INSERT INTO comments(user_id, movie_id, content) VALUES ($1, $2, $3) RETURNING id, user_id, movie_id, content, created_at
`

type AddCommentParams struct {
	UserID  int64  `json:"user_id"`
	MovieID int64  `json:"movie_id"`
	Content string `json:"content"`
}

func (q *Queries) AddComment(ctx context.Context, arg AddCommentParams) (Comment, error) {
	row := q.db.QueryRowContext(ctx, addComment, arg.UserID, arg.MovieID, arg.Content)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.MovieID,
		&i.Content,
		&i.CreatedAt,
	)
	return i, err
}

const getComments = `-- name: GetComments :many
SELECT comments.id, comments.content, comments.created_at, users.name, users.image
FROM comments 
INNER JOIN users ON users.id = comments.user_id
WHERE comments.movie_id = $1
`

type GetCommentsRow struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
}

func (q *Queries) GetComments(ctx context.Context, movieID int64) ([]GetCommentsRow, error) {
	rows, err := q.db.QueryContext(ctx, getComments, movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCommentsRow
	for rows.Next() {
		var i GetCommentsRow
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.CreatedAt,
			&i.Name,
			&i.Image,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
