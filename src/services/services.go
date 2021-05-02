package services

import (
	"github.com/felipeguilhermefs/selene/infra/config"
	"github.com/felipeguilhermefs/selene/repositories"
)

// NewServices init all services
func NewServices(cfg *config.Config, repos *repositories.Repositories) *Services {
	passwordService := newPasswordService(&cfg.Sec.Password)
	authService := newAuthService(repos.Session, repos.User, passwordService)

	return &Services{
		Auth:     authService,
		Password: passwordService,
	}
}

// Services all services in this app
type Services struct {
	Auth     AuthService
	Password PasswordService
}
