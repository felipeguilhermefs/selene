package config

import (
	"encoding/json"
	"os"
	"time"

	"github.com/felipeguilhermefs/selene/infra/database"
	"github.com/felipeguilhermefs/selene/infra/errors"
)

// Config hold all general configurable data of the server
type Config struct {
	Server struct {
		Port         int           `json:"port"`
		ReadTimeout  time.Duration `json:"read_timeout"`
		WriteTimeout time.Duration `json:"write_timeout"`
		IdleTimeout  time.Duration `json:"idle_timeout"`
	} `json:"server"`
	Postgres database.PostgresConfig `json:"postgres"`
	Sec      struct {
		Pepper string `json:"pepper"`
	} `json:"security"`
}

// ServerReadTimeout timeout to read entire request including the body
func (c *Config) ServerReadTimeout() time.Duration {
	return time.Second * c.Server.ReadTimeout
}

// ServerReadTimeout timeout to write the response
func (c *Config) ServerWriteTimeout() time.Duration {
	return time.Second * c.Server.WriteTimeout
}

// ServerReadTimeout idle connection timeout
func (c *Config) ServerIdleTimeout() time.Duration {
	return time.Second * c.Server.IdleTimeout
}

// LoadFromFile parse config from given file
func LoadFromFile(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "Opening "+filename)
	}
	defer f.Close()

	var cfg Config
	dec := json.NewDecoder(f)
	err = dec.Decode(&cfg)
	if err != nil {
		return nil, errors.Wrap(err, "Parsing "+filename)
	}

	return &cfg, nil
}
