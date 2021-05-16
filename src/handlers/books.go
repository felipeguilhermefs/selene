package handlers

import (
	"log"
	"net/http"

	"github.com/felipeguilhermefs/selene/context"
	"github.com/felipeguilhermefs/selene/services"
	"github.com/felipeguilhermefs/selene/view"
)

func HandleBooksPage(
	booksView *view.View,
	bookService services.BookService,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var vd view.Data

		user := context.User(r)

		books, err := bookService.GetBooks(user.ID)
		if err != nil {
			log.Println(err)
			booksView.Render(w, r, vd.WithError(err))
			return
		}

		booksView.Render(w, r, view.NewData(books))
	}
}
