package services

import (
	"github.com/felipeguilhermefs/selene/core/auth"
	"github.com/felipeguilhermefs/selene/infrastructure/config"
	"github.com/felipeguilhermefs/selene/infrastructure/session"
)

// Services all services in this app
type Services struct {
	Auth     AuthService
	Password PasswordService
}

// New init all services
func New(
	cfg config.ConfigStore,
	userRepository auth.UserRepository,
	sessionStore session.SessionStore,
) *Services {
	passwordService := newPasswordService(cfg)

	return &Services{
		Auth:     newAuthService(sessionStore, userRepository, passwordService),
		Password: passwordService,
	}
}
