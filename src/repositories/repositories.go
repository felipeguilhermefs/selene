package repositories

import (
	"github.com/gorilla/sessions"
	"gorm.io/gorm"

	"github.com/felipeguilhermefs/selene/models"
)

// NewRepositories init all repositories
func NewRepositories(db *gorm.DB, store sessions.Store) *Repositories {

	return &Repositories{
		db:      db,
		Session: newSessionRespository(store),
		User:    newUserRespository(db),
	}
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
