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
) auth.Handler {

	return func(w http.ResponseWriter, r *auth.Request) {
		var vd view.Data

		id, err := getUIntFromPath(r.Request, "id")
		if err != nil {
			bookView.Render(w, r.Request, vd.WithError(err))
		}

		book, err := bookService.GetBook(r.User.ID, id)
		if err != nil {
			bookView.Render(w, r.Request, vd.WithError(err))
		}

		form := bookForm{
			ID:       book.ID,
			Title:    book.Title,
			Author:   book.Author,
			Comments: book.Comments,
			Tags:     book.Tags,
		}

		bookView.Render(w, r.Request, view.NewData(&form))
	}
}

func HandleEditBook(
	bookView *view.View,
	bookService services.BookService,
) auth.Handler {

	return func(w http.ResponseWriter, r *auth.Request) {
		var form bookForm
		vd := view.NewData(&form)

		err := parseForm(r.Request, &form)
		if err != nil {
			bookView.Render(w, r.Request, vd.WithError(err))
			return
		}

		id, err := getUIntFromPath(r.Request, "id")
		if err != nil {
			bookView.Render(w, r.Request, vd.WithError(err))
		}

		book, err := bookService.GetBook(r.User.ID, id)
		if err != nil {
			bookView.Render(w, r.Request, vd.WithError(err))
		}

		book.Title = form.Title
		book.Author = form.Author
		book.Comments = form.Comments
		book.Tags = form.Tags

		if err := bookService.Update(book); err != nil {
			bookView.Render(w, r.Request, vd.WithError(err))
			return
		}

		http.Redirect(w, r.Request, "/books", http.StatusFound)
	}
}

func HandleDeleteBook(
	bookView *view.View,
	bookService services.BookService,
) auth.Handler {

	return func(w http.ResponseWriter, r *auth.Request) {
		var vd view.Data

		id, err := getUIntFromPath(r.Request, "id")
		if err != nil {
			bookView.Render(w, r.Request, vd.WithError(err))
		}

		err = bookService.Delete(r.User.ID, id)
		if err != nil {
			bookView.Render(w, r.Request, vd.WithError(err))
		}

		http.Redirect(w, r.Request, "/books", http.StatusFound)
	}
}
