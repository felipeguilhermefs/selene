package handlers

import (
	"log"
	"net/http"

	"github.com/felipeguilhermefs/selene/services"
	"github.com/felipeguilhermefs/selene/view"
)

type bookForm struct {
	Title    string `schema:"title"`
	Author   string `schema:"author"`
	Comments string `schema:"comments"`
	Tags     string `schema:"tags"`
}

func HandleNewBookPage(
	newBookView *view.View,
	authService services.AuthService,
	bookService services.BookService,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var vd view.Data

		_, err := authService.GetUser(r)
		if err != nil {
			log.Println(err)
			newBookView.Render(w, r, vd.WithError(err))
			return
		}

		var form bookForm
		parseURLParams(r, &form)
		newBookView.Render(w, r, view.NewData(&form))
	}
}
