package middlewares

import (
	"github.com/felipeguilhermefs/selene/infra/config"
	"github.com/gorilla/csrf"
)

func newCSRFMiddleware(cfg *config.SecurityConfig) Middleware {
	key := []byte(cfg.CSRF)
	return csrf.Protect(key, csrf.SameSite(csrf.SameSiteStrictMode))
}
