package services

import (
	"net/http"

	"github.com/gorilla/sessions"

	"github.com/felipeguilhermefs/selene/infra/config"
	"github.com/felipeguilhermefs/selene/infra/errors"
	"github.com/felipeguilhermefs/selene/models"
)

const (
	sessionCookie    = "session"
	authenticatedKey = "authenticated"
)

// SessionService handle operations over sessions
type SessionService interface {
	SignIn(w http.ResponseWriter, r *http.Request, user *models.User) error
	SignOut(w http.ResponseWriter, r *http.Request) error
	GetUser(r *http.Request) (*models.User, error)
}

// newSessionService creates a new instance of SessionService
func newSessionService(cfg *config.SessionConfig) SessionService {
	store := sessions.NewCookieStore(
		[]byte(cfg.AuthKey),
		[]byte(cfg.CryptoKey),
	)

	store.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
	}
	store.MaxAge(cfg.TTL)

	return &sessionService{
		store: store,
	}
}

type sessionService struct {
	store *sessions.CookieStore
}

func (ss *sessionService) SignIn(w http.ResponseWriter, r *http.Request, user *models.User) error {
	session, err := ss.store.Get(r, sessionCookie)
	if err != nil {
		return err
	}

	session.Values[authenticatedKey] = true

	return session.Save(r, w)
}

func (ss *sessionService) SignOut(w http.ResponseWriter, r *http.Request) error {
	session, err := ss.store.Get(r, sessionCookie)
	if err != nil {
		return err
	}

	session.Values[authenticatedKey] = false
	session.Options.MaxAge = -1

	return session.Save(r, w)
}

func (ss *sessionService) GetUser(r *http.Request) (*models.User, error) {
	session, err := ss.store.Get(r, sessionCookie)
	if err != nil {
		return nil, err
	}

	if auth, ok := session.Values[authenticatedKey].(bool); !ok || !auth {
		return nil, errors.ErrNotLoggedIn
	}

	return &models.User{Name: "Felipe"}, nil
}
