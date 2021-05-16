package middlewares

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/infra/config"
)

type Middleware = func(http.Handler) http.Handler

// Middlewares all middlewares in this app
type Middlewares struct {
	CSRF Middleware
}

// NewMiddlewares init all middlewares
func NewMiddlewares(cfg *config.Config) *Middlewares {
	return &Middlewares{
		CSRF: newCSRFMiddleware(&cfg.Sec),
	}
}
