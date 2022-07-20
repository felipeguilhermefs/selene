package repositories

import (
	"gorm.io/gorm"

	"github.com/felipeguilhermefs/selene/boundary"
)

// Repositories holds reference to all repositories
type Repositories struct {
	db *gorm.DB
}

// New init all repositories
func New(db *gorm.DB) *Repositories {

	return &Repositories{
		db: db,
	}
}

func (r *Repositories) allModels() []interface{} {
	return []interface{}{
		&boundary.User{},
		&boundary.Book{},
	}
}

// AutoMigrate updates all models in DB
func (r *Repositories) AutoMigrate() error {
	return r.db.AutoMigrate(r.allModels()...)
}
