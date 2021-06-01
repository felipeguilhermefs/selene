package middlewares

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/infra/config"
	"github.com/felipeguilhermefs/selene/services"
	"github.com/felipeguilhermefs/selene/view"
)

type Middleware = func(http.HandlerFunc) http.HandlerFunc

// Middlewares all middlewares in this app
type Middlewares struct {
	CSRF  Middleware
	CSP   Middleware
	Login Middleware
}

// NewMiddlewares init all middlewares
func NewMiddlewares(cfg *config.Config, authService services.AuthService, errorView *view.View) *Middlewares {
	return &Middlewares{
		CSRF:  newCSRFMiddleware(&cfg.Sec, errorView),
		CSP:   newCSPMiddleware(),
		Login: newLoginMiddleware(authService),
	}
}
