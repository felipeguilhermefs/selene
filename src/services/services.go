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

// NewServices init all services
func NewServices(cfg *config.Config, repos *repositories.Repositories) *Services {
	passwordService := newPasswordService(&cfg.Sec.Password)

	return &Services{
		Auth:     newAuthService(repos.Session, repos.User, passwordService),
		Password: passwordService,
		Book:     newBookService(repos.Book),
	}
}
