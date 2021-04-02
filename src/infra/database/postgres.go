package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/felipeguilhermefs/selene/infra/config"
	"github.com/felipeguilhermefs/selene/infra/errors"
)

// ConnectPostgres connect to DB and creates a connection pool
func ConnectPostgres(cfg *config.Postgres) (*gorm.DB, error) {
	pgDialect := postgres.Open(cfg.ConnectionInfo())

	db, err := gorm.Open(pgDialect, &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "Opening Connection")
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.Wrap(err, "Configuring Connection Pool")
	}

	fmt.Println(cfg)

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConnections())
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConnections())
	sqlDB.SetConnMaxLifetime(cfg.ConnectionTTL())

	return db, nil
}
