package handlers

import (
	"log"
	"net/http"

	"github.com/felipeguilhermefs/selene/services"
	"github.com/felipeguilhermefs/selene/view"
)

func HandleBooksPage(
	booksView *view.View,
	authService services.AuthService,
	bookService services.BookService,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var vd view.Data

		user, err := authService.GetUser(r)
		if err != nil {
			log.Println(err)
			booksView.Render(w, r, vd.WithError(err))
			return
		}

		books, err := bookService.GetBooks(user.ID)
		if err != nil {
			log.Println(err)
			booksView.Render(w, r, vd.WithError(err))
			return
		}

		booksView.Render(w, r, view.NewData(books))
	}
}
