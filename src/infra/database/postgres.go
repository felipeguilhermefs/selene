package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/felipeguilhermefs/selene/infra/config"
)

// ConnectPostgres connect to DB and creates a connection pool
func ConnectPostgres(cfg *config.DBConfig) (*gorm.DB, error) {
	pgDialect := postgres.Open(cfg.ConnInfo())

	db, err := gorm.Open(pgDialect, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn())
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn())
	sqlDB.SetConnMaxLifetime(cfg.ConnTTL())

	return db, nil
}
