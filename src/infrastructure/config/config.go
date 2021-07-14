package config

import "time"

type ConfigStore interface {
	Get(key, defaultValue string) string
	GetInt(key string, defaultValue int) int
	GetTime(key, defaultValue string) time.Duration
}

func New() ConfigStore {
	return &envConfigStore{}
}
