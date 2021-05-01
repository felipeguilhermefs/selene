package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/felipeguilhermefs/selene/infra/errors"
)

// PostgresConfig hold all configurable data needed to connect to postgres
type PostgresConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Conn     struct {
		MaxIdle    int           `json:"max_idle"`
		MaxOpen    int           `json:"max_open"`
		TTLMinutes time.Duration `json:"ttl_minutes"`
	} `json:"connection"`
}

// ConnectionInfo returns formated connection info to create a postgres connection
func (c *PostgresConfig) ConnectionInfo() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Name,
	)
}

// MaxIdleConnections returns max idle connections in postgresConfig pool
func (c *PostgresConfig) MaxIdleConnections() int {
	return c.Conn.MaxIdle
}

// MaxOpenConnections returns max open connections in postgresConfig pool
func (c *PostgresConfig) MaxOpenConnections() int {
	return c.Conn.MaxOpen
}

// ConnectionTTL connection lifetime in the pool
func (c *PostgresConfig) ConnectionTTL() time.Duration {
	return time.Minute * c.Conn.TTLMinutes
}

// ConnectPostgres connect to DB and creates a connection pool
func ConnectPostgres(cfg *PostgresConfig) (*gorm.DB, error) {
	pgDialect := postgres.Open(cfg.ConnectionInfo())

	db, err := gorm.Open(pgDialect, &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "Opening Connection")
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.Wrap(err, "Configuring Connection Pool")
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConnections())
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConnections())
	sqlDB.SetConnMaxLifetime(cfg.ConnectionTTL())

	return db, nil
}
