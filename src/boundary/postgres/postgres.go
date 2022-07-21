package postgres

import (
	"github.com/felipeguilhermefs/selene/infrastructure/config"
	"gorm.io/gorm"
)

func New(cfg config.ConfigStore) (*Postgres, error) {
	db, err := connect(cfg)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		db:             db,
		BookRepository: &PostgresBookRepository{db},
		UserRepository: &PostgresUserRepository{db},
	}, nil
}

type Postgres struct {
	db             *gorm.DB
	BookRepository *PostgresBookRepository
	UserRepository *PostgresUserRepository
}

func (pg *Postgres) RunMigrations() error {
	return pg.db.AutoMigrate(&Book{}, &User{})
}
