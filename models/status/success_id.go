package status

import (
	"net/http"

	"github.com/go-chi/render"
)

type SuccesWithID struct {
	StatusText string `json:"status"`
	ID         int64  `json:"id"`
}

func SuccessID(id int64, text string) render.Renderer {
	return &SuccesWithID{
		StatusText: text,
		ID:         id,
	}
}

func (d SuccesWithID) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, 200)
	return nil
}
