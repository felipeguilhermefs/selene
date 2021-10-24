package database

import (
	"time"

	"gorm.io/gorm"

	"github.com/felipeguilhermefs/selene/infrastructure/config"
)

const (
	maxConnectionRetries = 4
	backoffTime = 5 * time.Second
)

func Connect(cfg config.ConfigStore) (*gorm.DB, error) {
	db, err := connectPostgres(cfg)
	if err == nil {
		return db, nil	
	}

	// We try to connect to database a couple of times since 
	// it might not be available at startup time
	for try := 0; try < maxConnectionRetries; try++ {
		
		// A simple constant backoff should be ok for this simple case
		time.Sleep(backoffTime)
		
		db, err = connectPostgres(cfg)
		if err == nil {
			return db, nil	
		}
	}

	return nil, err
}
