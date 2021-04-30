package config

import (
	"encoding/json"
	"os"

	"github.com/felipeguilhermefs/selene/infra/database"
	"github.com/felipeguilhermefs/selene/infra/errors"
)

// Config hold all general configurable data of the server
type Config struct {
	Port     int                     `json:"port"`
	Postgres database.PostgresConfig `json:"postgres"`
	Sec      struct {
		Pepper string `json:"pepper"`
	} `json:"security"`
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
