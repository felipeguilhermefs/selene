package handlers

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/view"
)

func HandleMethodNotAllowed(methodNotAllowedView *view.View) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)

		var vd view.Data
		methodNotAllowedView.Render(w, r, &vd)
	}
}
