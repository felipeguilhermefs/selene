package auth

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/models"
)

type AuthService interface {
	GetUser(r *http.Request) (*models.User, error)
}
