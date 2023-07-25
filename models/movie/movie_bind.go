package movie

import (
	"errors"
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
	if len(mb.Title) < 3 {
		return errors.New("title can't be smaller than 3 char")
	}
	return nil
}
