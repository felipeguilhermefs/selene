package postgres

import (
	"github.com/felipeguilhermefs/selene/core/auth"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"not null;unique_index"`
	Password string `gorm:"not null"`
}

type PostgresUserRepository struct {
	db *gorm.DB
}

func (ur *PostgresUserRepository) Add(user *auth.NewUser) error {
	u := &User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	return ur.db.Create(u).Error
}

func (ur *PostgresUserRepository) FindOne(email string) (*auth.FullUser, error) {
	var record User

	err := ur.db.Where("email = ?", email).First(&record).Error
	if err == gorm.ErrRecordNotFound {
		return nil, auth.ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	return &auth.FullUser{
		ID:       record.ID,
		Name:     record.Name,
		Email:    record.Email,
		Password: record.Password,
	}, nil
}
