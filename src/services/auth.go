package services

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/infra/errors"
	"github.com/felipeguilhermefs/selene/models"
	"github.com/felipeguilhermefs/selene/repositories"
)

// AuthService handle operations over sessions
type AuthService interface {
	Login(w http.ResponseWriter, r *http.Request, email, password string) error
	SignUp(w http.ResponseWriter, r *http.Request, user *models.User) error
}

// newAuthService creates a new instance of AuthService
func newAuthService(
	sessionRepository repositories.SessionRepository,
	userRepository repositories.UserRepository,
	secretService SecretService,
) AuthService {

	return &authService{
		secretService:     secretService,
		sessionRepository: sessionRepository,
		userRepository:    userRepository,
	}
}

type authService struct {
	secretService     SecretService
	sessionRepository repositories.SessionRepository
	userRepository    repositories.UserRepository
}

func (as *authService) Login(w http.ResponseWriter, r *http.Request, email, password string) error {
	user, err := as.userRepository.ByEmail(email)
	if err != nil {
		return err
	}

	err = as.secretService.Compare(user.Secret, password)
	if err != nil {
		return err
	}

	return as.sessionRepository.SignIn(w, r, user)
}

func (as *authService) SignUp(w http.ResponseWriter, r *http.Request, user *models.User) error {
	if user.Email == "" {
		return errors.ErrEmailRequired
	}

	secret, err := as.secretService.Generate(user.Password)
	if err != nil {
		return err
	}
	user.Secret = secret
	user.Password = ""

	err = as.userRepository.Create(user)
	if err != nil {
		return err
	}

	return as.sessionRepository.SignIn(w, r, user)
}
