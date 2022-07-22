package handlers

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/core/auth"
	"github.com/felipeguilhermefs/selene/core/bookshelf"
	"github.com/felipeguilhermefs/selene/infrastructure/session"
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
	bookshelfControl *bookshelf.BookshelfControl,
	userVerifier auth.UserVerifier,
	sessionStore session.SessionStore,
) *Handlers {
	return &Handlers{
		SignupPage:  HandleSignupPage(&views.Signup),
		Signup:      HandleSignup(&views.Signup, srvcs.Auth),
		LoginPage:   HandleLoginPage(&views.Login),
		Login:       HandleLogin(&views.Login, userVerifier, sessionStore),
		Logout:      HandleLogout(sessionStore),
		BooksPage:   HandleBooksPage(&views.Books, bookshelfControl),
		NewBookPage: HandleNewBookPage(&views.NewBook),
		NewBook:     HandleNewBook(&views.NewBook, bookshelfControl),
		BookPage:    HandleBookPage(&views.EditBook, bookshelfControl),
		EditBook:    HandleEditBook(&views.EditBook, bookshelfControl),
		DeleteBook:  HandleDeleteBook(&views.EditBook, bookshelfControl),
		NotFound:    HandleNotFound(&views.NotFound),
		NotAuthentic: HandleError(
			&views.Error,
			http.StatusForbidden,
			"Sorry, authenticity check has failed.",
		),
	}
}
