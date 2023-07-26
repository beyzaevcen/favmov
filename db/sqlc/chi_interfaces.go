package db

import "net/http"

func (mv *Movie) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (mv *User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (em *EditMovieRow) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (em *GetCommentsRow) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (em *Comment) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (em *GetMyCommentsRow) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
