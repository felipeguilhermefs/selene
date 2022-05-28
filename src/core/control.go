package core

import (
	"strings"

	"github.com/felipeguilhermefs/selene/infra/errors"
)

type BookRepository interface {
	Insert(book *NewBook) error
}

type BookControl struct {
	BookRepository BookRepository
}

func (bc *BookControl) Add(book *NewBook) error {
	if book.UserID <= 0 {
		return errors.ErrUserIDRequired
	}

	if strings.TrimSpace(book.Title) == "" {
		return errors.ErrTitleRequired
	}

	if strings.TrimSpace(book.Author) == "" {
		return errors.ErrAuthorRequired
	}

	return bc.BookRepository.Insert(book)
}
