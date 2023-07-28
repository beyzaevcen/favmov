package main

import (
	"context"
	"database/sql"
	db "fav-mov/db/sqlc"
	"fav-mov/handler"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"google.golang.org/api/option"

	_ "github.com/lib/pq"
)

func main() {
	store := CreateStore()

	client := CreateFirebaseAuth()

	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(handler.ProvideStore(store), handler.ProvideFirebase(client))

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
		r.Post("/register", handler.RegisterUser)
		r.Post("/login", handler.LoginUser)
		r.With(handler.UserIDMiddleware).Get("/my", handler.GetUserById)
	})

	r.Route("/comments", func(r chi.Router) {
		r.Get("/{id}", handler.GetComments)

		r.Group(func(r chi.Router) {
			r.Use(handler.UserIDMiddleware)
			r.Post("/", handler.AddComment)
			r.Delete("/delete/{id}", handler.DeleteComment)
			r.Get("/my", handler.GetMyComments)
		})
	})

	r.Route("/watched_movie", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(handler.UserIDMiddleware)
			r.Get("/", handler.GetWatchedMovies)
			r.Post("/{movie_id}", handler.AddToWatchedMovies)
			r.Delete("/delete/{movie_id}", handler.DeleteFromWatchedMovies)

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

func CreateFirebaseAuth() *auth.Client {
	opt := option.WithCredentialsFile("favmov-53831-firebase-adminsdk-ji3yh-76fe49a863.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.TODO()

	context.Background()

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return client
}
