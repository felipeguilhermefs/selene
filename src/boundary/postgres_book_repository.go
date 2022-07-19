package boundary

import (
	"gorm.io/gorm"

	"github.com/felipeguilhermefs/selene/core/bookshelf"
	"github.com/felipeguilhermefs/selene/infra/errors"
	"github.com/felipeguilhermefs/selene/models"
)

type PostgresBookRepository struct {
	DB *gorm.DB
}

func (br *PostgresBookRepository) Insert(book *bookshelf.NewBook) error {
	b := &models.Book{
		Title:    book.Title,
		Author:   book.Author,
		Comments: book.Comments,
		Tags:     book.Tags,
		UserID:   book.UserID,
	}
	return br.DB.Create(b).Error
}

func (br *PostgresBookRepository) Update(book *bookshelf.UpdatedBook) error {
	b := &models.Book{
		Title:    book.Title,
		Author:   book.Author,
		Comments: book.Comments,
		Tags:     book.Tags,
	}
	return br.DB.Where("id = ?", book.ID).Updates(b).Error
}

func (br *PostgresBookRepository) FindOne(id uint) (*bookshelf.FullBook, error) {
	var record models.Book

	err := br.DB.Where("id = ?", id).First(&record).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.ErrBookNotFound
	}

	if err != nil {
		return nil, err
	}

	return &bookshelf.FullBook{
		ID:       record.ID,
		UserID:   record.UserID,
		Title:    record.Title,
		Author:   record.Author,
		Comments: record.Comments,
		Tags:     record.Tags,
	}, nil
}

func (br *PostgresBookRepository) FindMany(userID uint) ([]bookshelf.FullBook, error) {
	if userID <= 0 {
		return nil, errors.ErrUserIDRequired
	}

	var records []models.Book

	err := br.DB.Where("user_id = ?", userID).Find(&records).Error
	if err != nil {
		return []bookshelf.FullBook{}, err
	}

	books := make([]bookshelf.FullBook, len(records))
	for _, record := range records {
		book := bookshelf.FullBook{
			ID:       record.ID,
			UserID:   record.UserID,
			Title:    record.Title,
			Author:   record.Author,
			Comments: record.Comments,
			Tags:     record.Tags,
		}
		books = append(books, book)
	}
	return books, err
}

func (br *PostgresBookRepository) Delete(id uint) error {
	book := models.Book{Model: gorm.Model{ID: id}}
	return br.DB.Delete(&book).Error
}
