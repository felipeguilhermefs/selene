package services

import (
	"github.com/felipeguilhermefs/selene/infra/config"
	"github.com/felipeguilhermefs/selene/repositories"
)

// NewServices init all services
func NewServices(cfg *config.Config, repos *repositories.Repositories) *Services {
	secretSrvc := newSecretService(cfg)
	authSrvc := newAuthService(repos.Session, repos.User, secretSrvc)

	return &Services{
		Auth:    authSrvc,
		Secret:  secretSrvc,
	}
}

// Services all services in this app
type Services struct {
	Auth    AuthService
	Secret  SecretService
}
