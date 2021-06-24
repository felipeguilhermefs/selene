package handlers

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/services"
	"github.com/felipeguilhermefs/selene/view"
)

// Handlers all handlers in this app
type Handlers struct {
	SignupPage   http.HandlerFunc
	Signup       http.HandlerFunc
	LoginPage    http.HandlerFunc
	Login        http.HandlerFunc
	Logout       http.HandlerFunc
	BooksPage    http.HandlerFunc
	NewBookPage  http.HandlerFunc
	NewBook      http.HandlerFunc
	BookPage     http.HandlerFunc
	EditBook     http.HandlerFunc
	DeleteBook   http.HandlerFunc
	NotFound     http.HandlerFunc
	NotAuthentic http.HandlerFunc
	Error        http.HandlerFunc
}

// New init all handlers
func New(srvcs *services.Services, views *view.Views) *Handlers {
	return &Handlers{
		SignupPage:  HandleSignupPage(&views.Signup),
		Signup:      HandleSignup(&views.Signup, srvcs.Auth),
		LoginPage:   HandleLoginPage(&views.Login),
		Login:       HandleLogin(&views.Login, srvcs.Auth),
		Logout:      HandleLogout(srvcs.Auth),
		BooksPage:   HandleBooksPage(&views.Books, srvcs.Book),
		NewBookPage: HandleNewBookPage(&views.NewBook),
		NewBook:     HandleNewBook(&views.NewBook, srvcs.Book),
		BookPage:    HandleBookPage(&views.EditBook, srvcs.Book),
		EditBook:    HandleEditBook(&views.EditBook, srvcs.Book),
		DeleteBook:  HandleDeleteBook(&views.EditBook, srvcs.Book),
		NotFound:    HandleNotFound(&views.NotFound),
		NotAuthentic: HandleError(
			&views.Error,
			http.StatusForbidden,
			"Sorry, authenticity check has failed.",
		),
		Error: HandleError(
			&views.Error,
			http.StatusMethodNotAllowed,
			"Sorry, this HTTP method is not allowed in this route.",
		),
	}
}
