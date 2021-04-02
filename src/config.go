package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Config hold all general configurable data of the server
type Config struct {
	Port     int            `json:"port"`
	Postgres PostgresConfig `json:"postgres"`
}

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

// MaxIdleConnections returns max idle connections in postgres pool
func (c *PostgresConfig) MaxIdleConnections() int {
	return c.Conn.MaxIdle
}

// MaxOpenConnections returns max open connections in postgres pool
func (c *PostgresConfig) MaxOpenConnections() int {
	return c.Conn.MaxOpen
}

// ConnectionTTL connection lifetime in the pool
func (c *PostgresConfig) ConnectionTTL() time.Duration {
	return time.Minute * c.Conn.TTLMinutes
}

// LoadConfig loads config from file config.json
func LoadConfig() (*Config, error) {
	f, err := os.Open(".config")
	if err != nil {
		return nil, WrapError(err, "Opening .config")
	}
	defer f.Close()

	var cfg Config
	dec := json.NewDecoder(f)
	err = dec.Decode(&cfg)
	if err != nil {
		return nil, WrapError(err, "Parsing .config")
	}

	return &cfg, nil
}
