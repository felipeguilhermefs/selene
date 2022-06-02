package handlers

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/core"
	"github.com/felipeguilhermefs/selene/services"
	"github.com/felipeguilhermefs/selene/view"
)

// Handlers all handlers in this app
type Handlers struct {
	SignupPage   http.Handler
	Signup       http.Handler
	LoginPage    http.Handler
	Login        http.Handler
	Logout       http.Handler
	BooksPage    http.Handler
	NewBookPage  http.Handler
	NewBook      http.Handler
	BookPage     http.Handler
	EditBook     http.Handler
	DeleteBook   http.Handler
	NotFound     http.Handler
	NotAuthentic http.Handler
}

// New init all handlers
func New(
	srvcs *services.Services,
	views *view.Views,
	bookControl *core.BookControl,
) *Handlers {
	return &Handlers{
		SignupPage:  HandleSignupPage(&views.Signup),
		Signup:      HandleSignup(&views.Signup, srvcs.Auth),
		LoginPage:   HandleLoginPage(&views.Login),
		Login:       HandleLogin(&views.Login, srvcs.Auth),
		Logout:      HandleLogout(srvcs.Auth),
		BooksPage:   HandleBooksPage(&views.Books, srvcs.Book),
		NewBookPage: HandleNewBookPage(&views.NewBook),
		NewBook:     HandleNewBook(&views.NewBook, bookControl),
		BookPage:    HandleBookPage(&views.EditBook, bookControl),
		EditBook:    HandleEditBook(&views.EditBook, bookControl),
		DeleteBook:  HandleDeleteBook(&views.EditBook, bookControl),
		NotFound:    HandleNotFound(&views.NotFound),
		NotAuthentic: HandleError(
			&views.Error,
			http.StatusForbidden,
			"Sorry, authenticity check has failed.",
		),
	}
}
