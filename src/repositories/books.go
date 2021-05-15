package repositories

import (
	"strings"

	"gorm.io/gorm"

	"github.com/felipeguilhermefs/selene/infra/errors"
	"github.com/felipeguilhermefs/selene/models"
)

// BookRepository interacts with book DB
type BookRepository interface {
	Create(book *models.Book) error
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

func (br *bookRepository) Create(book *models.Book) error {
	if book.UserID <= 0 {
		return errors.ErrUserIDRequired
	}

	if strings.TrimSpace(book.Title) == "" {
		return errors.ErrTitleRequired
	}

	if strings.TrimSpace(book.Author) == "" {
		return errors.ErrAuthorRequired
	}

	return br.db.Create(book).Error
}

func (br *bookRepository) ByUserID(userID uint) ([]models.Book, error) {
	if userID <= 0 {
		return nil, errors.ErrUserIDRequired
	}

	var books []models.Book

	err := br.db.Where("user_id = ?", userID).Find(&books).Error
	return books, err
}
