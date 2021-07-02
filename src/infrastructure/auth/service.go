package auth

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/models"
)

type AuthService interface {
	Login(w http.ResponseWriter, r *http.Request, email, password string) error
	Logout(w http.ResponseWriter, r *http.Request) error
	SignUp(w http.ResponseWriter, r *http.Request, user *models.User) error
}
