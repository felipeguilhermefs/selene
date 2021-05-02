package handlers

import (
	"log"
	"net/http"

	"github.com/felipeguilhermefs/selene/services"
)

func HandleLogout(authService services.AuthService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		err := authService.Logout(w, r)
		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
