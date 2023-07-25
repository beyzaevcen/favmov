package movie

import (
	db "fav-mov/db/sqlc"
	"time"

	"net/http"
)

type MovieBind struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Score       float64   `json:"score"`
	Image       string    `json:"image"`
	ReleaseDate time.Time `json:"release_date"`
}

func (mb *MovieBind) AddMovieParams() db.AddMovieParams {
	return db.AddMovieParams{
		Title:       mb.Title,
		Description: mb.Description,
		Score:       mb.Score,
		Image:       mb.Image,
		ReleaseDate: mb.ReleaseDate,
	}
}

func (mb *MovieBind) Bind(r *http.Request) error {

	return nil
}

func (mb *MovieBind) ToUpdateMovieParams(id int64) db.EditMovieParams {
	return db.EditMovieParams{
		Description: mb.Description,
		Score:       mb.Score,
		Image:       mb.Image,
		ID:          id,
	}
}
