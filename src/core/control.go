package core

import (
	"strings"

	"github.com/felipeguilhermefs/selene/infra/errors"
)

type FullBook struct {
	ID       uint
	UserID   uint
	Title    string
	Author   string
	Comments string
	Tags     string
}

type BookRepository interface {
	Insert(book *NewBook) error
	Update(book *UpdatedBook) error
	Fetch(id uint) (*FullBook, error)
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

func (bc *BookControl) Update(book *UpdatedBook) error {
	if book.ID <= 0 {
		return errors.ErrIDInvalid
	}

	if book.UserID <= 0 {
		return errors.ErrUserIDRequired
	}

	if strings.TrimSpace(book.Title) == "" {
		return errors.ErrTitleRequired
	}

	if strings.TrimSpace(book.Author) == "" {
		return errors.ErrAuthorRequired
	}

	current, err := bc.BookRepository.Fetch(book.ID)
	if err != nil {
		return err
	}

	if current.UserID != book.UserID {
		return errors.ErrUserMismatch
	}

	return bc.BookRepository.Update(book)
}
