package database

import (
	"github.com/felipeguilhermefs/selene/infrastructure/config"
	"gorm.io/gorm"
)

func Connect(cfg config.ConfigStore) (*gorm.DB, error) {
	return connectPostgres(cfg)
}
