package services

import (
	"github.com/felipeguilhermefs/selene/models"
	"github.com/felipeguilhermefs/selene/repositories"
)

// BookService handle operations over sessions
type BookService interface {
	GetBooks(userID uint) ([]models.Book, error)
}

// newBookService creates a new instance of BookService
func newBookService(bookRepository repositories.BookRepository) BookService {

	return &bookService{
		bookRepository: bookRepository,
	}
}

type bookService struct {
	bookRepository repositories.BookRepository
}

func (bs *bookService) GetBooks(userID uint) ([]models.Book, error) {
	return bs.bookRepository.ByUserID(userID)
}