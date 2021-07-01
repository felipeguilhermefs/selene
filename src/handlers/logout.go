package handlers

import (
	"log"
	"net/http"

	"github.com/felipeguilhermefs/selene/infrastructure/auth"
)

func HandleLogout(authService auth.AuthService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		err := authService.Logout(w, r)
		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
