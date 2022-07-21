package postgres

import (
	"regexp"

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
		BookRepository: &PostgresBookRepository{DB: db},
		UserRepository: &PostgresUserRepository{
			DB:         db,
			EmailRegex: regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,16}$`),
		},
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
