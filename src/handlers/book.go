package handlers

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/infrastructure/auth"
	"github.com/felipeguilhermefs/selene/services"
	"github.com/felipeguilhermefs/selene/view"
)

func HandleBookPage(
	bookView *view.View,
	bookService services.BookService,
	authService auth.AuthService,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var vd view.Data

		id, err := getUIntFromPath(r, "id")
		if err != nil {
			bookView.Render(w, r, vd.WithError(err))
		}

		user, err := authService.GetUser(r)
		if err != nil {
			bookView.Render(w, r, vd.WithError(err))
		}

		book, err := bookService.GetBook(user.ID, id)
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
	bookService services.BookService,
	authService auth.AuthService,
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

		user, err := authService.GetUser(r)
		if err != nil {
			bookView.Render(w, r, vd.WithError(err))
		}

		book, err := bookService.GetBook(user.ID, id)
		if err != nil {
			bookView.Render(w, r, vd.WithError(err))
		}

		book.Title = form.Title
		book.Author = form.Author
		book.Comments = form.Comments
		book.Tags = form.Tags

		if err := bookService.Update(book); err != nil {
			bookView.Render(w, r, vd.WithError(err))
			return
		}

		http.Redirect(w, r, "/books", http.StatusFound)
	}
}

func HandleDeleteBook(
	bookView *view.View,
	bookService services.BookService,
	authService auth.AuthService,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var vd view.Data

		id, err := getUIntFromPath(r, "id")
		if err != nil {
			bookView.Render(w, r, vd.WithError(err))
		}

		user, err := authService.GetUser(r)
		if err != nil {
			bookView.Render(w, r, vd.WithError(err))
		}

		err = bookService.Delete(user.ID, id)
		if err != nil {
			bookView.Render(w, r, vd.WithError(err))
		}

		http.Redirect(w, r, "/books", http.StatusFound)
	}
}
