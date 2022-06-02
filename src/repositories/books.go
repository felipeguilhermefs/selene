package repositories

import (
	"gorm.io/gorm"

	"github.com/felipeguilhermefs/selene/infra/errors"
	"github.com/felipeguilhermefs/selene/models"
)

// BookRepository interacts with book DB
type BookRepository interface {
	ByUserAndID(userID, bookID uint) (*models.Book, error)
	ByUserID(userID uint) ([]models.Book, error)
}

// newBookRespository creates a new instance of newBookRespository
func newBookRespository(db *gorm.DB) BookRepository {

	return &bookRepository{
		db: db,
	}
}

type bookRepository struct {
	db *gorm.DB
}

func (br *bookRepository) ByUserAndID(userID, bookID uint) (*models.Book, error) {
	if userID <= 0 {
		return nil, errors.ErrUserIDRequired
	}

	if bookID <= 0 {
		return nil, errors.ErrIDInvalid
	}

	var book models.Book

	err := br.db.Where("user_id = ? AND id = ?", userID, bookID).First(&book).Error
	switch err {
	case gorm.ErrRecordNotFound:
		return nil, errors.ErrBookNotFound
	case nil:
		return &book, nil
	default:
		return nil, err
	}
}

func (br *bookRepository) ByUserID(userID uint) ([]models.Book, error) {
	if userID <= 0 {
		return nil, errors.ErrUserIDRequired
	}

	var books []models.Book

	err := br.db.Where("user_id = ?", userID).Find(&books).Error
	return books, err
}
