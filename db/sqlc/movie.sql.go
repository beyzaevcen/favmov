// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: movie.sql

package db

import (
	"context"
	"time"
)

const addMovie = `-- name: AddMovie :one
INSERT INTO movies (title, description, score, image, release_date) VALUES ($1,$2,$3,$4,$5) RETURNING id, title, description, score, image, release_date
`

type AddMovieParams struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Score       float64   `json:"score"`
	Image       string    `json:"image"`
	ReleaseDate time.Time `json:"release_date"`
}

func (q *Queries) AddMovie(ctx context.Context, arg AddMovieParams) (Movie, error) {
	row := q.db.QueryRowContext(ctx, addMovie,
		arg.Title,
		arg.Description,
		arg.Score,
		arg.Image,
		arg.ReleaseDate,
	)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Score,
		&i.Image,
		&i.ReleaseDate,
	)
	return i, err
}

const deleteMovie = `-- name: DeleteMovie :one
DELETE FROM movies 
WHERE id = $1 RETURNING id
`

func (q *Queries) DeleteMovie(ctx context.Context, id int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, deleteMovie, id)
	err := row.Scan(&id)
	return id, err
}

const editMovie = `-- name: EditMovie :one
UPDATE movies SET "description" = $1, "score" = $2, "image" = $3  
WHERE id = $4 RETURNING title, description, score, image, release_date
`

type EditMovieParams struct {
	Description string  `json:"description"`
	Score       float64 `json:"score"`
	Image       string  `json:"image"`
	ID          int64   `json:"id"`
}

type EditMovieRow struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Score       float64   `json:"score"`
	Image       string    `json:"image"`
	ReleaseDate time.Time `json:"release_date"`
}

func (q *Queries) EditMovie(ctx context.Context, arg EditMovieParams) (EditMovieRow, error) {
	row := q.db.QueryRowContext(ctx, editMovie,
		arg.Description,
		arg.Score,
		arg.Image,
		arg.ID,
	)
	var i EditMovieRow
	err := row.Scan(
		&i.Title,
		&i.Description,
		&i.Score,
		&i.Image,
		&i.ReleaseDate,
	)
	return i, err
}

const getMovie = `-- name: GetMovie :one
SELECT title, description, score, image, release_date 
FROM movies 
WHERE id =$1
`

type GetMovieRow struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Score       float64   `json:"score"`
	Image       string    `json:"image"`
	ReleaseDate time.Time `json:"release_date"`
}

func (q *Queries) GetMovie(ctx context.Context, id int64) (GetMovieRow, error) {
	row := q.db.QueryRowContext(ctx, getMovie, id)
	var i GetMovieRow
	err := row.Scan(
		&i.Title,
		&i.Description,
		&i.Score,
		&i.Image,
		&i.ReleaseDate,
	)
	return i, err
}

const getMovies = `-- name: GetMovies :many
SELECT id, title, description, score, image, release_date FROM movies
`

func (q *Queries) GetMovies(ctx context.Context) ([]Movie, error) {
	rows, err := q.db.QueryContext(ctx, getMovies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Movie
	for rows.Next() {
		var i Movie
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Score,
			&i.Image,
			&i.ReleaseDate,
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
