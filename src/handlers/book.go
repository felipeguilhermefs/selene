package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/felipeguilhermefs/selene/models"
	"github.com/felipeguilhermefs/selene/services"
	"github.com/felipeguilhermefs/selene/view"
	"github.com/gorilla/mux"
)

func HandleBookPage(
	bookView *view.View,
	authService services.AuthService,
	bookService services.BookService,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var vd view.Data

		user, err := authService.GetUser(r)
		if err != nil {
			log.Println(err)
			bookView.Render(w, r, vd.WithError(err))
			return
		}

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			log.Println(err)
			bookView.Render(w, r, vd.WithError(err))
		}

		book, err := bookService.GetBook(user.ID, uint(id))
		if err != nil {
			log.Println(err)
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

func HandleBook(
	bookView *view.View,
	authService services.AuthService,
	bookService services.BookService,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var form bookForm
		vd := view.NewData(&form)

		user, err := authService.GetUser(r)
		if err != nil {
			log.Println(err)
			bookView.Render(w, r, vd.WithError(err))
			return
		}

		err = parseForm(r, &form)
		if err != nil {
			log.Println(err)
			bookView.Render(w, r, vd.WithError(err))
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
			bookView.Render(w, r, vd.WithError(err))
			return
		}

		http.Redirect(w, r, "/books", http.StatusFound)
	}
}
