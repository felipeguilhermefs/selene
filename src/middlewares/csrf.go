package middlewares

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/infra/config"
	"github.com/gorilla/csrf"
)

func newCSRFMiddleware(cfg *config.SecurityConfig) Middleware {
	key := []byte(cfg.CSRF)

	csrfCheck := csrf.Protect(key, csrf.SameSite(csrf.SameSiteStrictMode))

	return func(next http.HandlerFunc) http.HandlerFunc {
		return csrfCheck(next).ServeHTTP
	}
}
