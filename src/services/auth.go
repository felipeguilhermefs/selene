package services

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/boundary/postgres"
	"github.com/felipeguilhermefs/selene/core/auth"
	"github.com/felipeguilhermefs/selene/infrastructure/session"
)

// AuthService handle operations over sessions
type AuthService interface {
	Login(w http.ResponseWriter, r *http.Request, email, password string) error
	Logout(w http.ResponseWriter, r *http.Request) error
	SignUp(w http.ResponseWriter, r *http.Request, user *postgres.User) error
}

// newAuthService creates a new instance of AuthService
func newAuthService(
	sessionStore session.SessionStore,
	userAdder auth.UserAdder,
	userVerifier auth.UserVerifier,
) AuthService {

	return &authService{
		sessionStore: sessionStore,
		userAdder:    userAdder,
		userVerifier: userVerifier,
	}
}

type authService struct {
	sessionStore session.SessionStore
	userAdder    auth.UserAdder
	userVerifier auth.UserVerifier
}

func (as *authService) Login(w http.ResponseWriter, r *http.Request, email, password string) error {
	user, err := as.userVerifier.Verify(email, password)
	if err != nil {
		return err
	}

	return as.sessionStore.SignIn(w, r, user.Email)
}

func (as *authService) Logout(w http.ResponseWriter, r *http.Request) error {
	return as.sessionStore.SignOut(w, r)
}

func (as *authService) SignUp(w http.ResponseWriter, r *http.Request, user *postgres.User) error {
	err := as.userAdder.Add(&auth.NewUser{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return err
	}

	return as.sessionStore.SignIn(w, r, user.Email)
}
