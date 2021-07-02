package auth

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/models"
)

type AuthenticatedRequest struct {
	Request *http.Request
	User    *models.User
}

type AuthenticatedHandler = func(w http.ResponseWriter, r *AuthenticatedRequest)

type Middleware = func(next AuthenticatedHandler) http.Handler

func NewMiddleware(authService AuthService) Middleware {

	return func(next AuthenticatedHandler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			user, err := authService.GetUser(r)
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}

			next(w, &AuthenticatedRequest{
				Request: r,
				User:    user,
			})
		})
	}
}
