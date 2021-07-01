package services

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/context"
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
) auth.AuthService {

	return &authService{
		passwordService:   passwordService,
		sessionRepository: sessionRepository,
		userRepository:    userRepository,
	}
}

type authService struct {
	passwordService   PasswordService
	sessionRepository repositories.SessionRepository
	userRepository    repositories.UserRepository
}

func (as *authService) GetUser(r *http.Request) (*models.User, error) {
	userCtx := r.Context().Value(context.UserKey)
	if userCtx != nil {
		if user, ok := userCtx.(*models.User); ok {
			return user, nil
		}
	}

	email, err := as.sessionRepository.GetUserEmail(r)
	if err != nil {
		return nil, err
	}

	user, err := as.userRepository.ByEmail(email)
	if err != nil {
		return nil, err
	}

	user.Secret = ""

	return user, nil
}

func (as *authService) Login(w http.ResponseWriter, r *http.Request, email, password string) error {
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

func (as *authService) Logout(w http.ResponseWriter, r *http.Request) error {
	return as.sessionRepository.SignOut(w, r)
}

func (as *authService) SignUp(w http.ResponseWriter, r *http.Request, user *models.User) error {
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
