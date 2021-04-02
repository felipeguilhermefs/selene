package config

import (
	"fmt"
	"time"
)

// Postgres hold all configurable data needed to connect to postgres
type Postgres struct {
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
func (c *Postgres) ConnectionInfo() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Name,
	)
}

// MaxIdleConnections returns max idle connections in postgres pool
func (c *Postgres) MaxIdleConnections() int {
	return c.Conn.MaxIdle
}

// MaxOpenConnections returns max open connections in postgres pool
func (c *Postgres) MaxOpenConnections() int {
	return c.Conn.MaxOpen
}

// ConnectionTTL connection lifetime in the pool
func (c *Postgres) ConnectionTTL() time.Duration {
	return time.Minute * c.Conn.TTLMinutes
}
