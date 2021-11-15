package csrf

import (
	"net/http"

	"github.com/gorilla/csrf"

	"github.com/felipeguilhermefs/selene/infrastructure/config"
)

func New(cfg config.ConfigStore) func(next http.Handler) http.Handler {
	secret := cfg.GetSecret("SELENE_CSRF_SECRET", "SecretWith32Chars...............")

	csrfCheck := csrf.Protect(
		[]byte(secret),
		csrf.SameSite(csrf.SameSiteStrictMode),
	)

	return func(next http.Handler) http.Handler {
		return csrfCheck(next)
	}
}
