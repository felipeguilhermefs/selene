package services

import (
	"github.com/felipeguilhermefs/selene/infrastructure/config"
	"github.com/felipeguilhermefs/selene/infrastructure/session"
	"github.com/felipeguilhermefs/selene/repositories"
)

// Services all services in this app
type Services struct {
	Auth     AuthService
	Password PasswordService
	Book     BookService
}

// New init all services
func New(
	cfg config.ConfigStore,
	repos *repositories.Repositories,
	sessionStore session.SessionStore,
) *Services {
	passwordService := newPasswordService(cfg)

	return &Services{
		Auth:     newAuthService(sessionStore, repos.User, passwordService),
		Password: passwordService,
		Book:     newBookService(repos.Book),
	}
}
