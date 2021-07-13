package session

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/infra/errors"
	"github.com/gorilla/sessions"
)

const (
	sessionCookie = "session"
	userKey       = "userID"
	noUser        = ""
)

func newCookieStore(cfg *Config) SessionStore {
	store := sessions.NewCookieStore(
		[]byte(cfg.AuthenticationKey),
		[]byte(cfg.EncryptionKey),
	)

	store.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
	store.MaxAge(cfg.TimeToLive)

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
