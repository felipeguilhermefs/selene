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
) auth.AuthenticatedHandler {

	return func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
		var vd view.Data

		books, err := bookService.GetBooks(r.User.ID)
		if err != nil {
			booksView.Render(w, r.Request, vd.WithError(err))
			return
		}

		booksView.Render(w, r.Request, view.NewData(books))
	}
}
