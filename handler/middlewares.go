package handler

import (
	"context"
	db "fav-mov/db/sqlc"
	"fav-mov/models/status"
	"net/http"
	"strconv"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/go-chi/render"
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

func UserIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		token := ctx.Value(FirebaseTokenKey).(*auth.Token)

		store := ctx.Value(StoreKey).(*db.Store)

		userID, err := store.GetUserIDByUID(ctx, token.UID)
		if err != nil {
			render.Render(w, r, status.ErrUnauthorized("user id could not be found"))
			return
		}

		ctx = context.WithValue(ctx, UserIdKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

func FirebaseAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		client := ctx.Value(FirebaseAuthKey).(*auth.Client)

		idToken := r.Header.Get("Authorization")
		splitToken := strings.Split(idToken, "Bearer ")

		if len(splitToken) != 2 {
			idToken = ""
		} else {
			idToken = splitToken[1]
		}

		if idToken == "" {
			render.Render(w, r, status.ErrUnauthorized("Missing authorization token."))
			return
		}

		token, err := client.VerifyIDToken(ctx, idToken)

		if err != nil {
			render.Render(w, r, status.ErrUnauthorized("Invalid token."))
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(ctx, FirebaseTokenKey, token)))
	})
}

func ProvideFirebase(client *auth.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), FirebaseAuthKey, client)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
