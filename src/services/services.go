package services

import (
	"github.com/felipeguilhermefs/selene/infra/config"
	"github.com/felipeguilhermefs/selene/repositories"
)

// NewServices init all services
func NewServices(cfg *config.Config, repos *repositories.Repositories) *Services {
	secretSrvc := newSecretService(cfg)
	userSrvc := newUserService(repos.User, secretSrvc)

	return &Services{
		User:   userSrvc,
		Secret: secretSrvc,
	}
}

// Services all services in this app
type Services struct {
	User   UserService
	Secret SecretService
}
