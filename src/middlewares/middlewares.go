package middlewares

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/services"
)

type Middleware = func(http.Handler) http.Handler

// Middlewares all middlewares in this app
type Middlewares struct {
	Login Middleware
}

// New init all middlewares
func New(authService services.AuthService) *Middlewares {
	return &Middlewares{
		Login: newLoginMiddleware(authService),
	}
}
