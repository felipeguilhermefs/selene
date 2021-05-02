package repositories

import (
	"github.com/gorilla/sessions"
	"gorm.io/gorm"

	"github.com/felipeguilhermefs/selene/infra/config"
	"github.com/felipeguilhermefs/selene/infra/database"
	"github.com/felipeguilhermefs/selene/infra/errors"
	"github.com/felipeguilhermefs/selene/models"
)

// NewRepositories init all repositories
func NewRepositories(cfg *config.Config) (*Repositories, error) {
	db, err := database.ConnectPostgres(&cfg.DB)
	if err != nil {
		return nil, errors.Wrap(err, "Connecting to Postgres")
	}

	store := sessions.NewCookieStore(
		[]byte(cfg.Sec.Session.AuthKey),
		[]byte(cfg.Sec.Session.CryptoKey),
	)

	store.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
	}
	store.MaxAge(cfg.Sec.Session.TTL)

	return &Repositories{
		db:      db,
		Session: newSessionRespository(store),
		User:    newUserRespository(db),
	}, nil
}

// Repositories holds reference to all repositories
type Repositories struct {
	db      *gorm.DB
	Session SessionRepository
	User    UserRepository
}

func (r *Repositories) allModels() []interface{} {
	return []interface{}{
		&models.User{},
	}
}

// AutoMigrate updates all models in DB
func (r *Repositories) AutoMigrate() error {
	return r.db.AutoMigrate(r.allModels()...)
}
