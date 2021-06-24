package middlewares

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/services"
)

type Middleware = func(http.HandlerFunc) http.HandlerFunc

// Middlewares all middlewares in this app
type Middlewares struct {
	CSRF       Middleware
	SecHeaders Middleware
	Login      Middleware
}

// New init all middlewares
func New(
	csrfSecret string,
	authService services.AuthService,
	notAuthentic http.HandlerFunc,
) *Middlewares {
	return &Middlewares{
		CSRF:       newCSRFMiddleware(csrfSecret, notAuthentic),
		SecHeaders: newSecHeaderMiddleware(),
		Login:      newLoginMiddleware(authService),
	}
}
