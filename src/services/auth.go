package services

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/infra/errors"
	"github.com/felipeguilhermefs/selene/infrastructure/auth"
	"github.com/felipeguilhermefs/selene/models"
	"github.com/felipeguilhermefs/selene/repositories"
)

// NewAuthService creates a new instance of AuthService
func NewAuthService(
	sessionRepository repositories.SessionRepository,
	userRepository repositories.UserRepository,
	passwordService PasswordService,
) *AuthService {

	return &AuthService{
		passwordService:   passwordService,
		sessionRepository: sessionRepository,
		userRepository:    userRepository,
	}
}

type AuthService struct {
	passwordService   PasswordService
	sessionRepository repositories.SessionRepository
	userRepository    repositories.UserRepository
}

func (as *AuthService) GetUser(r *http.Request) (*auth.User, error) {

	email, err := as.sessionRepository.GetUserEmail(r)
	if err != nil {
		return nil, err
	}

	user, err := as.userRepository.ByEmail(email)
	if err != nil {
		return nil, err
	}

	user.Secret = ""

	return &auth.User{
		ID: user.ID,
	}, nil
}

func (as *AuthService) Login(w http.ResponseWriter, r *http.Request, email, password string) error {
	user, err := as.userRepository.ByEmail(email)
	if err != nil {
		return err
	}

	err = as.passwordService.Compare(user.Secret, password)
	if err != nil {
		return err
	}

	return as.sessionRepository.SignIn(w, r, user)
}

func (as *AuthService) Logout(w http.ResponseWriter, r *http.Request) error {
	return as.sessionRepository.SignOut(w, r)
}

func (as *AuthService) SignUp(w http.ResponseWriter, r *http.Request, user *models.User) error {
	if user.Email == "" {
		return errors.ErrEmailRequired
	}

	secret, err := as.passwordService.Generate(user.Password)
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
