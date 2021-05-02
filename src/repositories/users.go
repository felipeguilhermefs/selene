package repositories

import (
	"regexp"
	"strings"

	"gorm.io/gorm"

	"github.com/felipeguilhermefs/selene/infra/errors"
	"github.com/felipeguilhermefs/selene/models"
)

const emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,16}$`

// UserRepository interacts with user DB
type UserRepository interface {
	Create(user *models.User) error
	ByEmail(email string) (*models.User, error)
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

func (ur *userRepository) Create(user *models.User) error {
	normalizedEmail := ur.normalizeEmail(user.Email)

	if !ur.emailRegex.MatchString(normalizedEmail) {
		return errors.ErrEmailInvalid
	}

	user.Email = normalizedEmail

	return ur.db.Create(user).Error
}

func (ur *userRepository) ByEmail(email string) (*models.User, error) {
	normalized := ur.normalizeEmail(email)

	if !ur.emailRegex.MatchString(normalized) {
		return nil, errors.ErrEmailInvalid
	}

	return ur.first("email = ?", normalized)
}

func (ur *userRepository) first(query interface{}, params ...interface{}) (*models.User, error) {
	var user models.User

	err := ur.db.Where(query, params...).First(&user).Error
	switch err {
	case gorm.ErrRecordNotFound:
		return nil, errors.ErrNotFound
	case nil:
		return &user, nil
	default:
		return nil, err
	}
}

func (ur *userRepository) normalizeEmail(input string) string {
	return strings.ToLower(strings.TrimSpace(input))
}
