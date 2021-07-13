package database

import (
	"fmt"
	"time"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	Conn     ConnConfig
}

type ConnConfig struct {
	MaxIdle int
	MaxOpen int
	TTL     time.Duration
}

// ConnInfo returns formated connection info to create a postgres connection
func (c *Config) ConnInfo() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Name,
	)
}
