package database

import "gorm.io/gorm"

func Connect(cfg *Config) (*gorm.DB, error) {
	return connectPostgres(cfg)
}
