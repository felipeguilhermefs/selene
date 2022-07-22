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
	userAdder auth.UserAdder,
	userFetcher auth.UserFetcher,
	sessionStore session.SessionStore,
	passwordEncripter auth.PasswordEncripter,
) *Services {
	passwordService := newPasswordService(cfg, passwordEncripter)

	return &Services{
		Auth:     newAuthService(sessionStore, userAdder, userFetcher, passwordService),
		Password: passwordService,
	}
}
