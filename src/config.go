package main

import (
	"encoding/json"
	"os"
)

// Config hold all general configurable data of the server
type Config struct {
	Port int `json:"port"`
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
