package handlers

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/context"
	"github.com/felipeguilhermefs/selene/core"
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
	bookAdder core.BookAdder,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var form bookForm
		vd := view.NewData(&form)

		err := parseForm(r, &form)
		if err != nil {
			newBookView.Render(w, r, vd.WithError(err))
			return
		}

		user := context.User(r)

		book := &core.NewBook{
			Title:    form.Title,
			Author:   form.Author,
			Comments: form.Comments,
			Tags:     form.Tags,
			UserID:   user.ID,
		}
		if err := bookAdder.Add(book); err != nil {
			newBookView.Render(w, r, vd.WithError(err))
			return
		}

		http.Redirect(w, r, "/books", http.StatusFound)
	}
}
