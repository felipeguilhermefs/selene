package handlers

import (
	"log"
	"net/http"

	"github.com/felipeguilhermefs/selene/infrastructure/auth"
)

func HandleLogout(authService auth.AuthService) auth.AuthenticatedHandler {

	return func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
		err := authService.Logout(w, r.Request)
		if err != nil {
			log.Println(err)
		}

		http.Redirect(w, r.Request, "/login", http.StatusFound)
	}
}
