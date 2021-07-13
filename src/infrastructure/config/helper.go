package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func getEnvTime(key, defaultValue string) (time.Duration, error) {
	value := os.Getenv(key)

	if value == "" {
		return time.ParseDuration(defaultValue)
	}

	res, err := time.ParseDuration(value)
	if err != nil {
		return res, fmt.Errorf("ENV %s: %w", key, err)
	}

	return res, nil
}

func getEnvInt(key string, defaultValue int) (int, error) {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue, nil
	}

	res, err := strconv.Atoi(value)
	if err != nil {
		return res, fmt.Errorf("ENV %s: %w", key, err)
	}

	return res, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
