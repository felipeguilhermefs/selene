package handlers

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/infrastructure/router"
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
	authenticated router.Middleware,
) *Handlers {
	return &Handlers{
		SignupPage:  HandleSignupPage(&views.Signup),
		Signup:      HandleSignup(&views.Signup, srvcs.Auth),
		LoginPage:   HandleLoginPage(&views.Login),
		Login:       HandleLogin(&views.Login, srvcs.Auth),
		Logout:      authenticated(HandleLogout(srvcs.Auth)),
		BooksPage:   authenticated(HandleBooksPage(&views.Books, srvcs.Book)),
		NewBookPage: authenticated(HandleNewBookPage(&views.NewBook)),
		NewBook:     authenticated(HandleNewBook(&views.NewBook, srvcs.Book)),
		BookPage:    authenticated(HandleBookPage(&views.EditBook, srvcs.Book)),
		EditBook:    authenticated(HandleEditBook(&views.EditBook, srvcs.Book)),
		DeleteBook:  authenticated(HandleDeleteBook(&views.EditBook, srvcs.Book)),
		NotFound:    HandleNotFound(&views.NotFound),
		NotAuthentic: HandleError(
			&views.Error,
			http.StatusForbidden,
			"Sorry, authenticity check has failed.",
		),
	}
}
