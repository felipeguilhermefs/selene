package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/felipeguilhermefs/selene/infra/config"
	"github.com/felipeguilhermefs/selene/infra/errors"
)

// ConnectPostgres connect to DB and creates a connection pool
func ConnectPostgres(cfg *config.DBConfig) (*gorm.DB, error) {
	pgDialect := postgres.Open(cfg.ConnInfo())

	db, err := gorm.Open(pgDialect, &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "Opening Connection")
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.Wrap(err, "Configuring Connection Pool")
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn())
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn())
	sqlDB.SetConnMaxLifetime(cfg.ConnTTL())

	return db, nil
}
