package main

import (
	"database/sql"
	db "fav-mov/db/sqlc"
	"fav-mov/handler"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	_ "github.com/lib/pq"
)

func main() {
	store := CreateStore()

	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(handler.ProvideStore(store))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Prenses başka bir kalede."))
	})

	r.Route("/movies", func(r chi.Router) {
		r.Get("/", handler.GetAllMovies)

		r.Group(func(r chi.Router) {
			r.Use(handler.AdminMiddleware)
			r.Post("/", handler.AddMovie)
			r.Delete("/delete/{id}", handler.DeleteMovie)
			r.Patch("/edit/{id}", handler.EditMovie)
		})
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", handler.GetAllUsers)
		r.Post("/", handler.RegisterUser)
	})

	r.Route("/comments", func(r chi.Router) {
		r.Get("/{id}", handler.GetComments)

		r.Group(func(r chi.Router) {
			r.Use(handler.IDMiddleware)
			r.Post("/", handler.AddComment)
			r.Delete("/delete/{id}", handler.DeleteComment)
			r.Get("/mycomments/{id}", handler.GetMyComments)
		})
	})

	http.ListenAndServe(":7770", r)

}

func CreateStore() *db.Store {
	database, err := sql.Open("postgres", "postgresql://adil:123456@localhost:5432/favmov?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db.NewStore(database)
}
