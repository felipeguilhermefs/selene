package repositories

import (
	"net/http"

	"github.com/gorilla/sessions"

	"github.com/felipeguilhermefs/selene/infra/errors"
	"github.com/felipeguilhermefs/selene/models"
)

const (
	sessionCookie      = "session"
	userKey            = "email"
	noUser             = ""
)

// SessionRepository interacts with session storage
type SessionRepository interface {
	GetUserEmail(r *http.Request) (string, error)
	SignIn(w http.ResponseWriter, r *http.Request, user *models.User) error
	SignOut(w http.ResponseWriter, r *http.Request) error
}

// newSessionRespository creates a new instance of SessionRepository
func newSessionRespository(store sessions.Store) SessionRepository {
	return &sessionRepository{store}
}

type sessionRepository struct {
	store sessions.Store
}

func (sr *sessionRepository) GetUserEmail(r *http.Request) (string, error) {
	session, err := sr.store.Get(r, sessionCookie)
	if err != nil {
		return "", err
	}

	email, ok := session.Values[userKey].(string)

	if !ok || email == noUser {
		return "", errors.ErrNotLoggedIn
	}

	return email, nil
}

func (sr *sessionRepository) SignIn(
	w http.ResponseWriter,
	r *http.Request,
	user *models.User,
) error {
	session, err := sr.store.Get(r, sessionCookie)
	if err != nil {
		return err
	}

	session.Values[userKey] = user.Email

	return session.Save(r, w)
}

func (sr *sessionRepository) SignOut(w http.ResponseWriter, r *http.Request) error {
	session, err := sr.store.Get(r, sessionCookie)
	if err != nil {
		return err
	}

	session.Values[userKey] = noUser
	session.Options.MaxAge = -1

	return session.Save(r, w)
}
