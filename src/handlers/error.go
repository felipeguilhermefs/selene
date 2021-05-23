package handlers

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/view"
)

func HandleError(
	errorView *view.View,
	statusCode int,
	errorMessage string,
) http.HandlerFunc {

	data := struct{ Message string }{errorMessage}

	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)

		errorView.Render(w, r, view.NewData(&data))
	}
}
