package repositories

import (
	"gorm.io/gorm"
	"github.com/felipeguilhermefs/selene/models"
)

// UserRepository interacts with user DB
type UserRepository interface {
	Create(user *models.User) error
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

