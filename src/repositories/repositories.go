package repositories

import (
	"gorm.io/gorm"

	"github.com/felipeguilhermefs/selene/models"
)

// Repositories holds reference to all repositories
type Repositories struct {
	db   *gorm.DB
	User UserRepository
	Book BookRepository
}

// New init all repositories
func New(db *gorm.DB) *Repositories {

	return &Repositories{
		db:   db,
		User: newUserRespository(db),
		Book: newBookRespository(db),
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
