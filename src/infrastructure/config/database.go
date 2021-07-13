package config

import "github.com/felipeguilhermefs/selene/infra/database"

// SELENE_DB_HOST: DB host.
// SELENE_DB_PORT: DB port.
// SELENE_DB_USER: DB user.
// SELENE_DB_PW: DB password.
// SELENE_DB_NAME: DB name.
// SELENE_DB_CONN_MAXIDLE: Max idle connections in the pool.
// SELENE_DB_CONN_MAXOPEN: Max open connections in the pool.
// SELENE_DB_CONN_TTL: Lifetime of a idle connection.
func loadDBConfig() (database.Config, error) {
	host := getEnv("SELENE_DB_HOST", "localhost")

	port, err := getEnvInt("SELENE_DB_PORT", 5432)
	if err != nil {
		return database.Config{}, err
	}

	user := getEnv("SELENE_DB_USER", "selene")
	password := getEnv("SELENE_DB_PW", "selene")
	dbName := getEnv("SELENE_DB_NAME", "selene")

	maxIdle, err := getEnvInt("SELENE_DB_CONN_MAXIDLE", 2)
	if err != nil {
		return database.Config{}, err
	}

	maxOpen, err := getEnvInt("SELENE_DB_CONN_MAXOPEN", 5)
	if err != nil {
		return database.Config{}, err
	}

	ttl, err := getEnvTime("SELENE_DB_CONN_TTL", "5m")
	if err != nil {
		return database.Config{}, err
	}

	return database.Config{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Name:     dbName,
		Conn: database.ConnConfig{
			MaxIdle: maxIdle,
			MaxOpen: maxOpen,
			TTL:     ttl,
		},
	}, nil
}
