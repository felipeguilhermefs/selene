package services

import (
	"github.com/felipeguilhermefs/selene/infra/config"
	"github.com/felipeguilhermefs/selene/repositories"
)

// NewServices init all services
func NewServices(cfg *config.Config, repos *repositories.Repositories) *Services {
	secretSrvc := newSecretService(cfg)
	sessionSrvc := newSessionService(&cfg.Sec.Session)
	userSrvc := newUserService(repos.User, secretSrvc)
	authSrvc := newAuthService(repos.Session, repos.User)

	return &Services{
		Auth:    authSrvc,
		User:    userSrvc,
		Session: sessionSrvc,
		Secret:  secretSrvc,
	}
}

// Services all services in this app
type Services struct {
	Auth    AuthService
	User    UserService
	Session SessionService
	Secret  SecretService
}
