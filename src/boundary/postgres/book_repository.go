package postgres

import (
	"gorm.io/gorm"

	"github.com/felipeguilhermefs/selene/core/bookshelf"
)

type Book struct {
	gorm.Model
	UserID   uint   `gorm:"not null;index"`
	Title    string `gorm:"not null"`
	Author   string `gorm:"not null"`
	Comments string
	Tags     string
}

type PostgresBookRepository struct {
	db *gorm.DB
}

func (br *PostgresBookRepository) Insert(book *bookshelf.NewBook) error {
	b := &Book{
		Title:    book.Title,
		Author:   book.Author,
		Comments: book.Comments,
		Tags:     book.Tags,
		UserID:   book.UserID,
	}
	return br.db.Create(b).Error
}

func (br *PostgresBookRepository) Update(book *bookshelf.UpdatedBook) error {
	b := &Book{
		Title:    book.Title,
		Author:   book.Author,
		Comments: book.Comments,
		Tags:     book.Tags,
	}
	return br.db.Where("id = ?", book.ID).Updates(b).Error
}

func (br *PostgresBookRepository) FindOne(id uint) (*bookshelf.FullBook, error) {
	var record Book

	err := br.db.Where("id = ?", id).First(&record).Error
	if err == gorm.ErrRecordNotFound {
		return nil, bookshelf.ErrBookNotFound
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
	var records []Book

	err := br.db.Where("user_id = ?", userID).Find(&records).Error
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
	book := Book{Model: gorm.Model{ID: id}}
	return br.db.Delete(&book).Error
}
