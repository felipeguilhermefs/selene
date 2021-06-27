package csrf

import (
	"net/http"

	"github.com/gorilla/csrf"
)

func New(secret string) func(next http.Handler) http.Handler {
	csrfCheck := csrf.Protect(
		[]byte(secret),
		csrf.SameSite(csrf.SameSiteStrictMode),
	)

	return func(next http.Handler) http.Handler {
		return csrfCheck(next)
	}
}
