package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectPostgres(cfg *Config) (*gorm.DB, error) {
	pgDialect := postgres.Open(buildConnString(cfg))

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

func buildConnString(cfg *Config) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name,
	)
}
