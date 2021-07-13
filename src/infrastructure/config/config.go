package config

import (
	"github.com/felipeguilhermefs/selene/infra/database"
	"github.com/felipeguilhermefs/selene/infrastructure/middleware/csrf"
	"github.com/felipeguilhermefs/selene/infrastructure/server"
	"github.com/felipeguilhermefs/selene/infrastructure/session"
	"github.com/felipeguilhermefs/selene/services"
)

type Config struct {
	CSRF     csrf.Config
	Password services.PasswordConfig
	Server   server.Config
	Session  session.Config
	DB       database.Config
}

func Load() (*Config, error) {
	server, err := loadServerConfig()
	if err != nil {
		return nil, err
	}

	session, err := loadSessionConfig()
	if err != nil {
		return nil, err
	}

	password, err := loadPasswordConfig()
	if err != nil {
		return nil, err
	}

	db, err := loadDBConfig()
	if err != nil {
		return nil, err
	}

	csrf := loadCSRFConfig()

	return &Config{
		CSRF:     csrf,
		DB:       db,
		Password: password,
		Server:   server,
		Session:  session,
	}, nil
}
