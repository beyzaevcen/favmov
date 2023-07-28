package handler

import (
	"context"
	db "fav-mov/db/sqlc"
	"fav-mov/models/status"
	"net/http"
	"strconv"

	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/lestrrat-go/jwx/jwt"
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
func ProvideJwtAuth(tokenAuth *jwtauth.JWTAuth) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), JwtAuthKey, tokenAuth)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func UserIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		token, claims, err := jwtauth.FromContext(ctx)

		if err != nil || token == nil || jwt.Validate(token) != nil {
			render.Render(w, r, status.ErrUnauthorized("Incorrect token."))
			return
		}

		userID, ok := claims["user_id"]
		if !ok {
			render.Render(w, r, status.ErrUnauthorized("Missing token"))
			return
		}

		ctx = context.WithValue(ctx, IDKey, int64(userID.(float64)))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
