package handlers

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/infrastructure/auth"
	"github.com/felipeguilhermefs/selene/models"
	"github.com/felipeguilhermefs/selene/services"
	"github.com/felipeguilhermefs/selene/view"
)

type bookForm struct {
	ID       uint   `schema:"id"`
	Title    string `schema:"title"`
	Author   string `schema:"author"`
	Comments string `schema:"comments"`
	Tags     string `schema:"tags"`
}

func HandleNewBookPage(newBookView *view.View) auth.Handler {

	return func(w http.ResponseWriter, r *auth.Request) {
		var form bookForm
		vd := view.NewData(&form)

		parseURLParams(r.Request, &form)
		newBookView.Render(w, r.Request, vd)
	}
}

func HandleNewBook(
	newBookView *view.View,
	bookService services.BookService,
) auth.Handler {

	return func(w http.ResponseWriter, r *auth.Request) {
		var form bookForm
		vd := view.NewData(&form)

		err := parseForm(r.Request, &form)
		if err != nil {
			newBookView.Render(w, r.Request, vd.WithError(err))
			return
		}

		book := &models.Book{
			Title:    form.Title,
			Author:   form.Author,
			Comments: form.Comments,
			Tags:     form.Tags,
			UserID:   r.User.ID,
		}
		if err := bookService.Create(book); err != nil {
			newBookView.Render(w, r.Request, vd.WithError(err))
			return
		}

		http.Redirect(w, r.Request, "/books", http.StatusFound)
	}
}
