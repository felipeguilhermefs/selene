package services

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/models"
	"github.com/felipeguilhermefs/selene/repositories"
)

// AuthService handle operations over sessions
type AuthService interface {
	SignUp(w http.ResponseWriter, r *http.Request, user *models.User) error
}

// newAuthService creates a new instance of AuthService
func newAuthService(
	sessionRepository repositories.SessionRepository,
	userRepository repositories.UserRepository,
) AuthService {

	return &authService{
		sessionRepository: sessionRepository,
		userRepository:    userRepository,
	}
}

type authService struct {
	sessionRepository repositories.SessionRepository
	userRepository    repositories.UserRepository
}

func (as *authService) SignUp(w http.ResponseWriter, r *http.Request, user *models.User) error {

	err := as.userRepository.Create(user)
	if err != nil {
		return err
	}

	return as.sessionRepository.SignIn(w, r, user)
}
