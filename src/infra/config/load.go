package config

import (
	"encoding/json"
	"os"
)

// LoadFromFile parse config from given file
func LoadFromFile(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	dec := json.NewDecoder(f)
	err = dec.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
