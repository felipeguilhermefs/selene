package postgres

import (
	"regexp"
	"strings"

	"github.com/felipeguilhermefs/selene/infra/errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"not null;unique_index"`
	Password string `gorm:"-"`
	Secret   string `gorm:"not null"`
}

type PostgresUserRepository struct {
	DB         *gorm.DB
	EmailRegex *regexp.Regexp
}

func (ur *PostgresUserRepository) Create(user *User) error {
	normalizedEmail := ur.normalizeEmail(user.Email)

	if !ur.EmailRegex.MatchString(normalizedEmail) {
		return errors.ErrEmailInvalid
	}

	user.Email = normalizedEmail

	return ur.DB.Create(user).Error
}

func (ur *PostgresUserRepository) ByEmail(email string) (*User, error) {
	normalized := ur.normalizeEmail(email)

	if !ur.EmailRegex.MatchString(normalized) {
		return nil, errors.ErrEmailInvalid
	}

	return ur.first("email = ?", normalized)
}

func (ur *PostgresUserRepository) first(query interface{}, params ...interface{}) (*User, error) {
	var user User

	err := ur.DB.Where(query, params...).First(&user).Error
	switch err {
	case gorm.ErrRecordNotFound:
		return nil, errors.ErrUserNotFound
	case nil:
		return &user, nil
	default:
		return nil, err
	}
}

func (ur *PostgresUserRepository) normalizeEmail(input string) string {
	return strings.ToLower(strings.TrimSpace(input))
}
