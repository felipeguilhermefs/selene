package repositories

import (
	"net/http"

	"github.com/gorilla/sessions"

	"github.com/felipeguilhermefs/selene/models"
)

const (
	sessionCookie = "session"
	idKey         = "id"
	emailKey      = "email"
)

// SessionRepository interacts with session storage
type SessionRepository interface {
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

func (sr *sessionRepository) SignIn(
	w http.ResponseWriter,
	r *http.Request,
	user *models.User,
) error {
	session, err := sr.store.Get(r, sessionCookie)
	if err != nil {
		return err
	}

	session.Values[idKey] = user.ID
	session.Values[emailKey] = user.Email

	return session.Save(r, w)
}

func (sr *sessionRepository) SignOut(w http.ResponseWriter, r *http.Request) error {
	session, err := sr.store.Get(r, sessionCookie)
	if err != nil {
		return err
	}

	session.Values[idKey] = -1
	session.Values[emailKey] = ""
	session.Options.MaxAge = -1

	return session.Save(r, w)
}
