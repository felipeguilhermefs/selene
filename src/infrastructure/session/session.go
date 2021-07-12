package session

import "net/http"

type UserID = string

// SessionStore interacts with session storage
type SessionStore interface {
	GetUserID(r *http.Request) (UserID, error)
	SignIn(w http.ResponseWriter, r *http.Request, userID UserID) error
	SignOut(w http.ResponseWriter, r *http.Request) error
}
