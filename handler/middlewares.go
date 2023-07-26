package handler

import (
	"context"
	db "fav-mov/db/sqlc"
	"net/http"
	"strconv"
)

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("AdminKey")
		ids, err := strconv.ParseInt(id, 10, 64)
		if err != nil || id != "0" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("ID int olmayabilir"))
			return
		}
		ctx := context.WithValue(r.Context(), AdminKey, ids)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ProvideStore(store *db.Store) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), StoreKey, store)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func IDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("UserId")
		ids, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("ID int olmayabilir"))
			return
		}

		ctx := context.WithValue(r.Context(), IDKey, ids)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
