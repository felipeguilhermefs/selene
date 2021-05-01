package config

import "time"

// ServerConfig server config data
type ServerConfig struct {
	Port     int           `json:"port"`
	RTimeout time.Duration `json:"read_timeout"`
	WTimeout time.Duration `json:"write_timeout"`
	ITimeout time.Duration `json:"idle_timeout"`
}

// ReadTimeout timeout for reading an entire request, including body
func (c *ServerConfig) ReadTimeout() time.Duration {
	return c.RTimeout * defaultTimeUnit
}

// WriteTimeout timeout for writing response
func (c *ServerConfig) WriteTimeout() time.Duration {
	return c.WTimeout * defaultTimeUnit
}

// IdleTimeout connection timeout receiving keep alives
func (c *ServerConfig) IdleTimeout() time.Duration {
	return c.ITimeout * defaultTimeUnit
}
