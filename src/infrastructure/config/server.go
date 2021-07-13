package config

import "github.com/felipeguilhermefs/selene/infrastructure/server"

// SELENE_SERVER_PORT: server listening port.
// SELENE_SERVER_READ_TIMEOUT: timeout for reading an entire request, including body.
// SELENE_SERVER_WRITE_TIMEOUT: timeout for writing response.
// SELENE_SERVER_IDLE_TIMEOUT: connection timeout receiving keep alives.
func loadServerConfig() (server.Config, error) {
	port, err := getEnvInt("SELENE_SERVER_PORT", 8000)
	if err != nil {
		return server.Config{}, err
	}

	readTimeout, err := getEnvTime("SELENE_SERVER_READ_TIMEOUT", "15s")
	if err != nil {
		return server.Config{}, err
	}

	writeTimeout, err := getEnvTime("SELENE_SERVER_WRITE_TIMEOUT", "15s")
	if err != nil {
		return server.Config{}, err
	}

	idleTimeout, err := getEnvTime("SELENE_SERVER_IDLE_TIMEOUT", "60s")
	if err != nil {
		return server.Config{}, err
	}

	return server.Config{
		Port:         port,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}, nil
}
