package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/felipeguilhermefs/selene/infra/errors"
)

// ConnectDatabase initiates DB abstraction
// multiple a connection pool will be created
func ConnectDatabase(dbConfig *PostgresConfig) (*gorm.DB, error) {
	pgDialect := postgres.Open(dbConfig.ConnectionInfo())

	db, err := gorm.Open(pgDialect, &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "Opening DB Connection")
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.Wrap(err, "Configuring Connection Pool")
	}

	fmt.Println(dbConfig)

	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConnections())
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConnections())
	sqlDB.SetConnMaxLifetime(dbConfig.ConnectionTTL())

	return db, nil
}
