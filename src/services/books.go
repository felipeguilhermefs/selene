package services

import (
	"github.com/felipeguilhermefs/selene/models"
	"github.com/felipeguilhermefs/selene/repositories"
)

// BookService handle operations over sessions
type BookService interface {
	Create(book *models.Book) error
	Update(book *models.Book) error
	GetBook(userID, bookID uint) (*models.Book, error)
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

func (bs *bookService) Create(book *models.Book) error {
	return bs.bookRepository.Create(book)
}

func (bs *bookService) Update(book *models.Book) error {
	return bs.bookRepository.Update(book)
}

func (bs *bookService) GetBook(userID, bookID uint) (*models.Book, error) {
	return bs.bookRepository.ByUserAndID(userID, bookID)
}

func (bs *bookService) GetBooks(userID uint) ([]models.Book, error) {
	return bs.bookRepository.ByUserID(userID)
}
