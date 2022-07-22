package services

import (
	"github.com/felipeguilhermefs/selene/core/auth"
	"github.com/felipeguilhermefs/selene/infrastructure/config"
	"github.com/felipeguilhermefs/selene/infrastructure/session"
)

// Services all services in this app
type Services struct {
	Auth AuthService
}

// New init all services
func New(
	cfg config.ConfigStore,
	userAdder auth.UserAdder,
	userVerifier auth.UserVerifier,
	sessionStore session.SessionStore,
	passwordComparer auth.PasswordComparer,
) *Services {
	return &Services{
		Auth: newAuthService(sessionStore, userAdder, userVerifier),
	}
}
