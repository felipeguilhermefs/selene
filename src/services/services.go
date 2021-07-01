package services

import (
	"github.com/felipeguilhermefs/selene/infra/config"
	"github.com/felipeguilhermefs/selene/repositories"
)

// Services all services in this app
type Services struct {
	Password PasswordService
	Book     BookService
}

// New init all services
func New(passwordConfig *config.PasswordConfig, repos *repositories.Repositories) *Services {
	passwordService := newPasswordService(passwordConfig)

	return &Services{
		Password: passwordService,
		Book:     newBookService(repos.Book),
	}
}
