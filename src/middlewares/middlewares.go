package middlewares

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/infra/config"
	"github.com/felipeguilhermefs/selene/services"
)

type Middleware = func(http.HandlerFunc) http.HandlerFunc

// Middlewares all middlewares in this app
type Middlewares struct {
	CSRF  Middleware
	Login Middleware
}

// NewMiddlewares init all middlewares
func NewMiddlewares(cfg *config.Config, authService services.AuthService) *Middlewares {
	return &Middlewares{
		CSRF:  newCSRFMiddleware(&cfg.Sec),
		Login: newLoginMiddleware(authService),
	}
}
