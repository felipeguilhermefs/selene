package services

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/boundary/postgres"
	"github.com/felipeguilhermefs/selene/core/auth"
	"github.com/felipeguilhermefs/selene/infra/errors"
	"github.com/felipeguilhermefs/selene/infrastructure/session"
)

// AuthService handle operations over sessions
type AuthService interface {
	GetUser(r *http.Request) (*auth.FullUser, error)
	Login(w http.ResponseWriter, r *http.Request, email, password string) error
	Logout(w http.ResponseWriter, r *http.Request) error
	SignUp(w http.ResponseWriter, r *http.Request, user *postgres.User) error
}

// newAuthService creates a new instance of AuthService
func newAuthService(
	sessionStore session.SessionStore,
	userControl *auth.UserControl,
	passwordService PasswordService,
) AuthService {

	return &authService{
		passwordService: passwordService,
		sessionStore:    sessionStore,
		userControl:     userControl,
	}
}

type authService struct {
	passwordService PasswordService
	sessionStore    session.SessionStore
	userControl     *auth.UserControl
}

func (as *authService) GetUser(r *http.Request) (*auth.FullUser, error) {
	email, err := as.sessionStore.GetUserID(r)
	if err != nil {
		return nil, err
	}

	user, err := as.userControl.FetchOne(email)
	if err != nil {
		return nil, err
	}

	user.Password = ""

	return user, nil
}

func (as *authService) Login(w http.ResponseWriter, r *http.Request, email, password string) error {
	user, err := as.userControl.FetchOne(email)
	if err != nil {
		return err
	}

	err = as.passwordService.Compare(user.Password, password)
	if err != nil {
		return err
	}

	return as.sessionStore.SignIn(w, r, user.Email)
}

func (as *authService) Logout(w http.ResponseWriter, r *http.Request) error {
	return as.sessionStore.SignOut(w, r)
}

func (as *authService) SignUp(w http.ResponseWriter, r *http.Request, user *postgres.User) error {
	if user.Email == "" {
		return errors.ErrEmailRequired
	}

	secret, err := as.passwordService.Generate(user.Password)
	if err != nil {
		return err
	}

	err = as.userControl.Add(&auth.NewUser{
		Name:     user.Name,
		Email:    user.Email,
		Password: secret,
	})
	if err != nil {
		return err
	}

	return as.sessionStore.SignIn(w, r, user.Email)
}
