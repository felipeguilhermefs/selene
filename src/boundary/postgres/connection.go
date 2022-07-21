package postgres

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/felipeguilhermefs/selene/infrastructure/config"
)

const (
	maxConnectionRetries = 4
	backoffTime          = 5 * time.Second
)

func connect(cfg config.ConfigStore) (*gorm.DB, error) {
	connString := buildConnString(cfg)

	db, err := tryConnect(connString, 0)
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

func tryConnect(connString string, retry int) (*gorm.DB, error) {
	pgDialect := postgres.Open(connString)

	db, err := gorm.Open(pgDialect, &gorm.Config{})
	if err != nil {
		// We try to connect to database a couple of times since
		// it might not be available at startup time
		if retry < maxConnectionRetries {
			// A simple constant backoff should be ok for this simple case
			time.Sleep(backoffTime)

			return tryConnect(connString, retry+1)
		}

		return nil, err
	}

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
