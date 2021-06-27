package middleware

import (
	"net/http"

	"github.com/gorilla/csrf"
)

func NewCSRF(secret string) Middleware {
	csrfCheck := csrf.Protect(
		[]byte(secret),
		csrf.SameSite(csrf.SameSiteStrictMode),
	)

	return func(next http.Handler) http.Handler {
		return csrfCheck(next)
	}
}
