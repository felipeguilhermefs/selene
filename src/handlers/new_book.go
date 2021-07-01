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

func HandleNewBookPage(newBookView *view.View) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var form bookForm
		vd := view.NewData(&form)

		parseURLParams(r, &form)
		newBookView.Render(w, r, vd)
	}
}

func HandleNewBook(
	newBookView *view.View,
	bookService services.BookService,
	authService auth.AuthService,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var form bookForm
		vd := view.NewData(&form)

		err := parseForm(r, &form)
		if err != nil {
			newBookView.Render(w, r, vd.WithError(err))
			return
		}

		user, err := authService.GetUser(r)
		if err != nil {
			newBookView.Render(w, r, vd.WithError(err))
		}

		book := &models.Book{
			Title:    form.Title,
			Author:   form.Author,
			Comments: form.Comments,
			Tags:     form.Tags,
			UserID:   user.ID,
		}
		if err := bookService.Create(book); err != nil {
			newBookView.Render(w, r, vd.WithError(err))
			return
		}

		http.Redirect(w, r, "/books", http.StatusFound)
	}
}
