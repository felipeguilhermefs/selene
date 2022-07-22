package handlers

import (
	"log"
	"net/http"

	"github.com/felipeguilhermefs/selene/infrastructure/session"
)

func HandleLogout(sessionStore session.SessionStore) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		err := sessionStore.SignOut(w, r)
		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
