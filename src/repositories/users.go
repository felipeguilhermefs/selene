package repositories

import (
	"gorm.io/gorm"

	"github.com/felipeguilhermefs/selene/infra/errors"
	"github.com/felipeguilhermefs/selene/models"
)

// UserRepository interacts with user DB
type UserRepository interface {
	Create(user *models.User) error
	ByEmail(email string) (*models.User, error)
}

// newUserRespository creates a new instance of UserRepository
func newUserRespository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

type userRepository struct {
	db *gorm.DB
}

func (ur *userRepository) Create(user *models.User) error {
	return ur.db.Create(user).Error
}

func (ur *userRepository) ByEmail(email string) (*models.User, error) {
	return ur.first("email = ?", email)
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
