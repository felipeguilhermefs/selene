package boundary

import (
	"gorm.io/gorm"

	"github.com/felipeguilhermefs/selene/core"
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
