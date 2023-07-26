package handler

import (
	db "fav-mov/db/sqlc"
	"fav-mov/models/movie"
	"fav-mov/models/status"
	"fav-mov/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
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

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	store := ctx.Value(StoreKey).(*db.Store)

	movieId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		render.Render(w, r, status.ErrBadRequest(err))
		return
	}

	id, err := store.DeleteMovie(ctx, movieId)
	if err != nil {
		render.Render(w, r, status.ErrNotFoundOrInternal(err))
		return
	}

	render.Render(w, r, status.SuccessID(id, "Successfully deleted :)"))
}
func EditMovie(w http.ResponseWriter, r *http.Request) {
	movieParams := &movie.MovieBind{}
	err := render.Bind(r, movieParams)
	if err != nil {
		render.Render(w, r, status.ErrBadRequest(err))
		return
	}

	ctx := r.Context()
	store := ctx.Value(StoreKey).(*db.Store)

	movieID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		render.Render(w, r, status.ErrBadRequest(err))
		return
	}

	mb, err := store.EditMovie(ctx, movieParams.ToUpdateMovieParams(movieID))
	if err != nil {
		render.Render(w, r, status.ErrNotFoundOrInternal(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, &mb)
}
