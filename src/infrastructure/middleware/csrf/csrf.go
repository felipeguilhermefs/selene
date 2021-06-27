package csrf

import (
	"net/http"

	"github.com/gorilla/csrf"
)

type Config interface {
	Secret() string
}

func New(cfg Config) func(next http.Handler) http.Handler {
	csrfCheck := csrf.Protect(
		[]byte(cfg.Secret()),
		csrf.SameSite(csrf.SameSiteStrictMode),
	)

	return func(next http.Handler) http.Handler {
		return csrfCheck(next)
	}
}
