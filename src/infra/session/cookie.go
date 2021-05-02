package session

import (
	"github.com/gorilla/sessions"

	"github.com/felipeguilhermefs/selene/infra/config"
)

// NewCookieStore init a cookie session store
func NewCookieStore(cfg *config.SessionConfig) sessions.Store {
	cookieStore := sessions.NewCookieStore(
		[]byte(cfg.AuthKey),
		[]byte(cfg.CryptoKey),
	)

	cookieStore.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
	}
	cookieStore.MaxAge(cfg.TTL)

	return cookieStore
}
