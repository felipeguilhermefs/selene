package handlers

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/context"
	"github.com/felipeguilhermefs/selene/core/auth"
	"github.com/felipeguilhermefs/selene/core/bookshelf"
	"github.com/felipeguilhermefs/selene/view"
)

func HandleBookPage(
	bookView *view.View,
	bookFetcher bookshelf.BookFetcher,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var vd view.Data

		id, err := getUIntFromPath(r, "id")
		if err != nil {
			bookView.Render(w, r, vd.WithError(err))
		}

		user := context.User(r)
		if user == nil {
			bookView.Render(w, r, vd.WithError(auth.ErrNotLoggedIn))
		}

		book, err := bookFetcher.FetchOne(user.ID, id)
		if err != nil {
			bookView.Render(w, r, vd.WithError(err))
		}

		form := bookForm{
			ID:       book.ID,
			Title:    book.Title,
			Author:   book.Author,
			Comments: book.Comments,
			Tags:     book.Tags,
		}

		bookView.Render(w, r, view.NewData(&form))
	}
}

func HandleEditBook(
	bookView *view.View,
	bookUpdater bookshelf.BookUpdater,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var form bookForm
		vd := view.NewData(&form)

		err := parseForm(r, &form)
		if err != nil {
			bookView.Render(w, r, vd.WithError(err))
			return
		}

		id, err := getUIntFromPath(r, "id")
		if err != nil {
			bookView.Render(w, r, vd.WithError(err))
		}

		user := context.User(r)
		if user == nil {
			bookView.Render(w, r, vd.WithError(auth.ErrNotLoggedIn))
		}

		book := &bookshelf.UpdatedBook{
			ID:       id,
			UserID:   user.ID,
			Title:    form.Title,
			Author:   form.Author,
			Comments: form.Comments,
			Tags:     form.Tags,
		}

		if err := bookUpdater.Update(book); err != nil {
			bookView.Render(w, r, vd.WithError(err))
			return
		}

		http.Redirect(w, r, "/books", http.StatusFound)
	}
}

func HandleDeleteBook(
	bookView *view.View,
	bookRemover bookshelf.BookRemover,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var vd view.Data

		id, err := getUIntFromPath(r, "id")
		if err != nil {
			bookView.Render(w, r, vd.WithError(err))
		}

		user := context.User(r)

		err = bookRemover.Remove(user.ID, id)
		if err != nil {
			bookView.Render(w, r, vd.WithError(err))
		}

		http.Redirect(w, r, "/books", http.StatusFound)
	}
}
