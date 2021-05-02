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
