package db

import "net/http"

func (mv *Movie) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (em *EditMovieRow) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
