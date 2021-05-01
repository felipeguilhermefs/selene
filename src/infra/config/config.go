package config

import (
	"fmt"
	"time"
)

const timeUnit = time.Second

// Config hold all general configurable data of the server
type Config struct {
	Server struct {
		Port         int           `json:"port"`
		ReadTimeout  time.Duration `json:"read_timeout"`
		WriteTimeout time.Duration `json:"write_timeout"`
		IdleTimeout  time.Duration `json:"idle_timeout"`
	} `json:"server"`
	PG struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Conn     struct {
			MaxIdle int           `json:"max_idle"`
			MaxOpen int           `json:"max_open"`
			TTL     time.Duration `json:"ttl"`
		} `json:"connection"`
	} `json:"postgres"`
	Sec struct {
		Pepper string `json:"pepper"`
	} `json:"security"`
}

// ServerReadTimeout timeout to read entire request including the body
func (c *Config) ServerReadTimeout() time.Duration {
	return timeUnit * c.Server.ReadTimeout
}

// ServerReadTimeout timeout to write the response
func (c *Config) ServerWriteTimeout() time.Duration {
	return timeUnit * c.Server.WriteTimeout
}

// ServerReadTimeout idle connection timeout
func (c *Config) ServerIdleTimeout() time.Duration {
	return timeUnit * c.Server.IdleTimeout
}

// PGConnInfo returns formated connection info to create a postgres connection
func (c *Config) PGConnInfo() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.PG.Host, c.PG.Port, c.PG.User, c.PG.Password, c.PG.Name,
	)
}

// PGConnTTL postgres connection lifetime in the pool
func (c *Config) PGConnTTL() time.Duration {
	return timeUnit * c.PG.Conn.TTL
}
