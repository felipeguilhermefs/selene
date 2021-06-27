package config

import "time"

// ServerConfig server config data
type ServerConfig struct {
	Prt      int           `json:"port"`
	RTimeout time.Duration `json:"read_timeout"`
	WTimeout time.Duration `json:"write_timeout"`
	ITimeout time.Duration `json:"idle_timeout"`
}

func (c *ServerConfig) Port() int {
	return c.Prt
}

func (c *ServerConfig) ReadTimeout() time.Duration {
	return c.RTimeout * defaultTimeUnit
}

func (c *ServerConfig) WriteTimeout() time.Duration {
	return c.WTimeout * defaultTimeUnit
}

func (c *ServerConfig) IdleTimeout() time.Duration {
	return c.ITimeout * defaultTimeUnit
}
