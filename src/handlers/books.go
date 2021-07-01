package handlers

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/infrastructure/auth"
	"github.com/felipeguilhermefs/selene/services"
	"github.com/felipeguilhermefs/selene/view"
)

func HandleBooksPage(
	booksView *view.View,
	bookService services.BookService,
	authService auth.AuthService,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var vd view.Data

		user, err := authService.GetUser(r)
		if err != nil {
			booksView.Render(w, r, vd.WithError(err))
		}

		books, err := bookService.GetBooks(user.ID)
		if err != nil {
			booksView.Render(w, r, vd.WithError(err))
			return
		}

		booksView.Render(w, r, view.NewData(books))
	}
}
