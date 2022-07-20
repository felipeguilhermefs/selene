package repositories

import (
	"regexp"
	"strings"

	"gorm.io/gorm"

	"github.com/felipeguilhermefs/selene/boundary"
	"github.com/felipeguilhermefs/selene/infra/errors"
)

const emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,16}$`

// UserRepository interacts with user DB
type UserRepository interface {
	Create(user *boundary.User) error
	ByEmail(email string) (*boundary.User, error)
}

// newUserRespository creates a new instance of UserRepository
func newUserRespository(db *gorm.DB) UserRepository {

	return &userRepository{
		db:         db,
		emailRegex: regexp.MustCompile(emailRegex),
	}
}

type userRepository struct {
	db         *gorm.DB
	emailRegex *regexp.Regexp
}

func (ur *userRepository) Create(user *boundary.User) error {
	normalizedEmail := ur.normalizeEmail(user.Email)

	if !ur.emailRegex.MatchString(normalizedEmail) {
		return errors.ErrEmailInvalid
	}

	user.Email = normalizedEmail

	return ur.db.Create(user).Error
}

func (ur *userRepository) ByEmail(email string) (*boundary.User, error) {
	normalized := ur.normalizeEmail(email)

	if !ur.emailRegex.MatchString(normalized) {
		return nil, errors.ErrEmailInvalid
	}

	return ur.first("email = ?", normalized)
}

func (ur *userRepository) first(query interface{}, params ...interface{}) (*boundary.User, error) {
	var user boundary.User

	err := ur.db.Where(query, params...).First(&user).Error
	switch err {
	case gorm.ErrRecordNotFound:
		return nil, errors.ErrUserNotFound
	case nil:
		return &user, nil
	default:
		return nil, err
	}
}

func (ur *userRepository) normalizeEmail(input string) string {
	return strings.ToLower(strings.TrimSpace(input))
}
