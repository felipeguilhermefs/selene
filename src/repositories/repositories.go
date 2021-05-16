package repositories

import (
	"github.com/gorilla/sessions"
	"gorm.io/gorm"

	"github.com/felipeguilhermefs/selene/models"
)

// Repositories holds reference to all repositories
type Repositories struct {
	db      *gorm.DB
	Session SessionRepository
	User    UserRepository
	Book    BookRepository
}

// NewRepositories init all repositories
func NewRepositories(db *gorm.DB, store sessions.Store) *Repositories {

	return &Repositories{
		db:      db,
		Session: newSessionRespository(store),
		User:    newUserRespository(db),
		Book:    newBookRespository(db),
	}
}

func (r *Repositories) allModels() []interface{} {
	return []interface{}{
		&models.User{},
		&models.Book{},
	}
}

// AutoMigrate updates all models in DB
func (r *Repositories) AutoMigrate() error {
	return r.db.AutoMigrate(r.allModels()...)
}
