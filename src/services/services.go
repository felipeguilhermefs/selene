package services

import (
	"github.com/felipeguilhermefs/selene/infra/config"
	"github.com/felipeguilhermefs/selene/repositories"
)

// Services all services in this app
type Services struct {
	Auth     AuthService
	Password PasswordService
	Book     BookService
}

// New init all services
func New(passwordConfig *config.PasswordConfig, repos *repositories.Repositories) *Services {
	passwordService := newPasswordService(passwordConfig)

	return &Services{
		Auth:     newAuthService(repos.Session, repos.User, passwordService),
		Password: passwordService,
		Book:     newBookService(repos.Book),
	}
}
