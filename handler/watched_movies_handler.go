package handler

import (
	db "fav-mov/db/sqlc"
	"fav-mov/models/status"
	"fav-mov/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func GetWatchedMovies(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	store := ctx.Value(StoreKey).(*db.Store)
	userId := ctx.Value(IDKey).(int64)

	movies, err := store.GetWatchedMovies(ctx, userId)
	if err != nil {
		render.Render(w, r, status.ErrInternal(err))
		return
	}

	moviesPayload := utils.Map(movies, func(el db.Movie) render.Renderer { return &el })

	render.Status(r, http.StatusOK)
	render.RenderList(w, r, utils.NewRenderList(moviesPayload))

}

func AddToWatchedMovies(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	store := ctx.Value(StoreKey).(*db.Store)
	userId := ctx.Value(IDKey).(int64)

	movieId, err := strconv.ParseInt(chi.URLParam(r, "movie_id"), 10, 64)
	if err != nil {
		render.Render(w, r, status.ErrBadRequest(err))
		return
	}

	movie, err := store.AddToWatchedMovies(ctx, db.AddToWatchedMoviesParams{
		UserID:  userId,
		MovieID: movieId,
	})
	if err != nil {
		render.Render(w, r, status.ErrInternal(err))
		return
	}
	movieInfo, err := store.GetMovie(ctx, movieId)
	if err != nil {
		render.Render(w, r, status.ErrInternal(err))
		return
	}

	movieRow := db.Movie{
		ID:          movie.MovieID,
		Title:       movieInfo.Title,
		Description: movieInfo.Description,
		Score:       movieInfo.Score,
		Image:       movieInfo.Image,
		ReleaseDate: movieInfo.ReleaseDate,
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, &movieRow)
}

func DeleteFromWatchedMovies(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	store := ctx.Value(StoreKey).(*db.Store)
	userId := ctx.Value(IDKey).(int64)

	movieId, err := strconv.ParseInt(chi.URLParam(r, "movie_id"), 10, 64)
	if err != nil {
		render.Render(w, r, status.ErrBadRequest(err))
		return
	}

	movie, err := store.DeleteFromWatchedMovies(ctx, db.DeleteFromWatchedMoviesParams{
		UserID:  userId,
		MovieID: movieId,
	})
	if err != nil {
		render.Render(w, r, status.ErrNotFoundOrInternal(err))
		return
	}

	render.Render(w, r, status.SuccessID(movie.MovieID, "Succesfully deleted :)"))
}
