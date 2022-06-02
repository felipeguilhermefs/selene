package handlers

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/context"
	"github.com/felipeguilhermefs/selene/core"
	"github.com/felipeguilhermefs/selene/view"
)

func HandleBooksPage(
	booksView *view.View,
	bookFetcher core.BookFetcher,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var vd view.Data

		user := context.User(r)

		books, err := bookFetcher.FetchMany(user.ID)
		if err != nil {
			booksView.Render(w, r, vd.WithError(err))
			return
		}

		booksView.Render(w, r, view.NewData(books))
	}
}
