package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

type envConfigStore struct{}

func (ecs *envConfigStore) Get(key, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

func (ecs *envConfigStore) GetInt(key string, defaultValue int) int {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	res, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Error converting env var %s (%s) falling back to default %d", key, err, defaultValue)
		return defaultValue
	}

	return res
}

func (ecs *envConfigStore) GetTime(key, defaultValue string) time.Duration {
	value := os.Getenv(key)

	if value == "" {
		value = defaultValue
	}

	res, err := time.ParseDuration(value)
	if err != nil {
		log.Printf("Error converting env var %s (%s) falling back to 0s", key, err)
		return time.Duration(0)
	}

	return res
}
