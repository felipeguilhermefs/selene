package handlers

import (
	"log"
	"net/http"

	"github.com/felipeguilhermefs/selene/models"
	"github.com/felipeguilhermefs/selene/services"
	"github.com/felipeguilhermefs/selene/view"
)

type bookForm struct {
	Title    string `schema:"title"`
	Author   string `schema:"author"`
	Comments string `schema:"comments"`
	Tags     string `schema:"tags"`
}

func HandleNewBookPage(
	newBookView *view.View,
	authService services.AuthService,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var form bookForm
		vd := view.NewData(&form)

		_, err := authService.GetUser(r)
		if err != nil {
			log.Println(err)
			newBookView.Render(w, r, vd.WithError(err))
			return
		}

		parseURLParams(r, &form)
		newBookView.Render(w, r, view.NewData(&form))
	}
}

func HandleNewBook(
	newBookView *view.View,
	authService services.AuthService,
	bookService services.BookService,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var form bookForm
		vd := view.NewData(&form)

		user, err := authService.GetUser(r)
		if err != nil {
			log.Println(err)
			newBookView.Render(w, r, vd.WithError(err))
			return
		}

		err = parseForm(r, &form)
		if err != nil {
			log.Println(err)
			newBookView.Render(w, r, vd.WithError(err))
			return
		}

		book := &models.Book{
			Title:    form.Title,
			Author:   form.Author,
			Comments: form.Comments,
			Tags:     form.Tags,
			UserID:   user.ID,
		}
		if err := bookService.Create(book); err != nil {
			log.Println(err)
			newBookView.Render(w, r, vd.WithError(err))
			return
		}

		http.Redirect(w, r, "/books", http.StatusFound)
	}
}
