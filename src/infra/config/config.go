package config

import "time"

const defaultTimeUnit = time.Second

// Config hold all general configurable data of the server
type Config struct {
	Server ServerConfig   `json:"server"`
	DB     DBConfig       `json:"database"`
	Sec    SecurityConfig `json:"security"`
}
