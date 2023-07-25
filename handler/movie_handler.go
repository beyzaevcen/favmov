package handler

import (
	db "fav-mov/db/sqlc"
	"fav-mov/models/movie"
	"fav-mov/models/status"
	"fav-mov/utils"
	"net/http"

	"github.com/go-chi/render"
)

func AddMovie(w http.ResponseWriter, r *http.Request) {
	movieParams := &movie.MovieBind{}
	err := render.Bind(r, movieParams)
	if err != nil {
		render.Render(w, r, status.ErrBadRequest(err))
		return
	}

	ctx := r.Context()
	store := ctx.Value(StoreKey).(*db.Store)

	movie, err := store.AddMovie(ctx, movieParams.AddMovieParams())
	if err != nil {
		render.Render(w, r, status.ErrInternal(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, &movie)
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	store := ctx.Value(StoreKey).(*db.Store)

	movies, err := store.GetMovies(ctx)
	if err != nil {
		render.Render(w, r, status.ErrInternal(err))
		return
	}

	moviesPayload := utils.Map(movies, func(el db.Movie) render.Renderer { return &el })

	render.Status(r, http.StatusOK)
	render.RenderList(w, r, utils.NewRenderList(moviesPayload))
}
