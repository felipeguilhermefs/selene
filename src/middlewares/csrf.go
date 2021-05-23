package middlewares

import (
	"net/http"

	"github.com/gorilla/csrf"

	"github.com/felipeguilhermefs/selene/handlers"
	"github.com/felipeguilhermefs/selene/infra/config"
	"github.com/felipeguilhermefs/selene/view"
)

func newCSRFMiddleware(cfg *config.SecurityConfig, errorView *view.View) Middleware {
	notAuthentic := handlers.HandleError(
		errorView,
		http.StatusForbidden,
		"Sorry, authenticity check has failed.",
	)

	csrfCheck := csrf.Protect(
		[]byte(cfg.CSRF),
		csrf.SameSite(csrf.SameSiteStrictMode),
		csrf.ErrorHandler(notAuthentic),
	)

	return func(next http.HandlerFunc) http.HandlerFunc {
		return csrfCheck(next).ServeHTTP
	}
}
