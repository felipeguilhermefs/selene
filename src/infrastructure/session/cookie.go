package session

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/infra/errors"
	"github.com/felipeguilhermefs/selene/infrastructure/config"
	"github.com/gorilla/sessions"
)

const (
	sessionCookie = "session"
	userKey       = "userID"
	noUser        = ""
)

func newCookieStore(cfg config.ConfigStore) SessionStore {
	store := sessions.NewCookieStore(
		[]byte(cfg.Get("SELENE_SESSION_AUTH_KEY", "AuthKeyWith64Chars..............................................")),
		[]byte(cfg.Get("SELENE_SESSION_CRYPTO_KEY", "CryptoKeyWith32Chars............")),
	)

	store.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
	store.MaxAge(cfg.GetInt("SELENE_SESSION_TTL", 900))

	return &cookieStore{store}
}

type cookieStore struct {
	store sessions.Store
}

func (ss *cookieStore) GetUserID(r *http.Request) (UserID, error) {
	session, err := ss.store.Get(r, sessionCookie)
	if err != nil {
		return "", err
	}

	userID, ok := session.Values[userKey].(UserID)

	if !ok || userID == noUser {
		return "", errors.ErrNotLoggedIn
	}

	return userID, nil
}

func (sr *cookieStore) SignIn(w http.ResponseWriter, r *http.Request, userID UserID) error {
	session, err := sr.store.Get(r, sessionCookie)
	if err != nil {
		return err
	}

	session.Values[userKey] = userID

	return session.Save(r, w)
}

func (sr *cookieStore) SignOut(w http.ResponseWriter, r *http.Request) error {
	session, err := sr.store.Get(r, sessionCookie)
	if err != nil {
		return err
	}

	session.Values[userKey] = noUser
	session.Options.MaxAge = -1

	return session.Save(r, w)
}
