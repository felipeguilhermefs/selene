package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectPostgres connect to DB and creates a connection pool
func ConnectPostgres(cfg *Config) (*gorm.DB, error) {
	pgDialect := postgres.Open(cfg.ConnInfo())

	db, err := gorm.Open(pgDialect, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(cfg.Conn.MaxIdle)
	sqlDB.SetMaxOpenConns(cfg.Conn.MaxOpen)
	sqlDB.SetConnMaxLifetime(cfg.Conn.TTL)

	return db, nil
}
