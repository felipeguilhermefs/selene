package session

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/infrastructure/config"
)

type UserID = string

type SessionStore interface {
	GetUserID(r *http.Request) (UserID, error)
	SignIn(w http.ResponseWriter, r *http.Request, userID UserID) error
	SignOut(w http.ResponseWriter, r *http.Request) error
}

func NewStore(cfg config.ConfigStore) SessionStore {
	return newCookieStore(cfg)
}
