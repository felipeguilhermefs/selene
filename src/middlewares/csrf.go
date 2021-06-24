package middlewares

import (
	"net/http"

	"github.com/gorilla/csrf"
)

func newCSRFMiddleware(secret string, notAuthentic http.HandlerFunc) Middleware {
	csrfCheck := csrf.Protect(
		[]byte(secret),
		csrf.SameSite(csrf.SameSiteStrictMode),
		csrf.ErrorHandler(notAuthentic),
	)

	return func(next http.HandlerFunc) http.HandlerFunc {
		return csrfCheck(next).ServeHTTP
	}
}
