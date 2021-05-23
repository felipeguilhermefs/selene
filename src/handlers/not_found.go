package handlers

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/view"
)

func HandleNotFound(notFoundView *view.View) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)

		var vd view.Data
		notFoundView.Render(w, r, &vd)
	}
}
