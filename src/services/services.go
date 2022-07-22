package services

import (
	"github.com/felipeguilhermefs/selene/core/auth"
	"github.com/felipeguilhermefs/selene/infrastructure/session"
)

// Services all services in this app
type Services struct {
	Auth AuthService
}

// New init all services
func New(
	userAdder auth.UserAdder,
	sessionStore session.SessionStore,
) *Services {
	return &Services{
		Auth: newAuthService(sessionStore, userAdder),
	}
}
