package boundary

import (
	"gorm.io/gorm"

	"github.com/felipeguilhermefs/selene/core"
	"github.com/felipeguilhermefs/selene/infra/errors"
	"github.com/felipeguilhermefs/selene/models"
)

type PostgresBookRepository struct {
	DB *gorm.DB
}

func (br *PostgresBookRepository) Insert(book *core.NewBook) error {
	b := &models.Book{
		Title:    book.Title,
		Author:   book.Author,
		Comments: book.Comments,
		Tags:     book.Tags,
		UserID:   book.UserID,
	}
	return br.DB.Create(b).Error
}

func (br *PostgresBookRepository) Update(book *core.UpdatedBook) error {
	b := &models.Book{
		Title:    book.Title,
		Author:   book.Author,
		Comments: book.Comments,
		Tags:     book.Tags,
	}
	return br.DB.Model(&models.Book{}).Where("id = ?", book.ID).Updates(b).Error
}

func (br *PostgresBookRepository) Fetch(id uint) (*core.FullBook, error) {
	var book models.Book

	err := br.DB.Where("id = ?", id).First(&book).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.ErrBookNotFound
	}

	if err != nil {
		return nil, err
	}

	return &core.FullBook{
		ID:       book.ID,
		UserID:   book.UserID,
		Title:    book.Title,
		Author:   book.Author,
		Comments: book.Comments,
		Tags:     book.Tags,
	}, nil
}

func (br *PostgresBookRepository) Delete(id uint) error {
	book := models.Book{Model: gorm.Model{ID: id}}
	return br.DB.Delete(&book).Error
}
