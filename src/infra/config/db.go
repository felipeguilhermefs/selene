package config

import (
	"fmt"
	"time"
)

// DBConfig database config data
type DBConfig struct {
	Host     string       `json:"host"`
	Port     int          `json:"port"`
	User     string       `json:"user"`
	Password string       `json:"password"`
	Name     string       `json:"name"`
	Conn     DBConnConfig `json:"connection"`
}

// DBConnConfig
type DBConnConfig struct {
	MaxIdle int           `json:"max_idle"`
	MaxOpen int           `json:"max_open"`
	TTL     time.Duration `json:"ttl"`
}

// ConnInfo returns formated connection info to create a postgres connection
func (c *DBConfig) ConnInfo() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Name,
	)
}

// MaxIdleConn max number of idle connections
func (c *DBConfig) MaxIdleConn() int {
	return c.Conn.MaxIdle
}

// MaxOpenConn max number of open connections
func (c *DBConfig) MaxOpenConn() int {
	return c.Conn.MaxOpen
}

// ConnTTL max lifetime of a connection to be reused
func (c *DBConfig) ConnTTL() time.Duration {
	return c.Conn.TTL * defaultTimeUnit
}
