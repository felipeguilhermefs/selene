package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/felipeguilhermefs/selene/infrastructure/config"
)

func connectPostgres(cfg config.ConfigStore) (*gorm.DB, error) {
	pgDialect := postgres.Open(buildConnString(cfg))

	db, err := gorm.Open(pgDialect, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(cfg.GetInt("SELENE_DB_CONN_MAXIDLE", 2))
	sqlDB.SetMaxOpenConns(cfg.GetInt("SELENE_DB_CONN_MAXOPEN", 5))
	sqlDB.SetConnMaxLifetime(cfg.GetTime("SELENE_DB_CONN_TTL", "5m"))

	return db, nil
}

func buildConnString(cfg config.ConfigStore) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Get("SELENE_DB_HOST", "localhost"),
		cfg.GetInt("SELENE_DB_PORT", 5432),
		cfg.GetSecret("SELENE_DB_USER", "selene"),
		cfg.GetSecret("SELENE_DB_PW", "selene"),
		cfg.GetSecret("SELENE_DB_NAME", "selene"),
	)
}
