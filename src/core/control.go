package core

import (
	"strings"

	"github.com/felipeguilhermefs/selene/infra/errors"
)

type BookRepository interface {
	Insert(book *NewBook) error
	Update(book *UpdatedBook) error
	FindOne(id uint) (*FullBook, error)
	FindMany(userID uint) ([]FullBook, error)
	Delete(id uint) error
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

	current, err := bc.BookRepository.FindOne(book.ID)
	if err != nil {
		return err
	}

	if current.UserID != book.UserID {
		return errors.ErrUserMismatch
	}

	return bc.BookRepository.Update(book)
}

func (bc *BookControl) Remove(userID uint, id uint) error {
	if id <= 0 {
		return errors.ErrIDInvalid
	}

	if userID <= 0 {
		return errors.ErrUserIDRequired
	}

	book, err := bc.BookRepository.FindOne(id)
	if err != nil {
		return err
	}

	if book.UserID != userID {
		return errors.ErrUserMismatch
	}

	return bc.BookRepository.Delete(id)
}

func (bc *BookControl) FetchOne(userID, id uint) (*FullBook, error) {
	if id <= 0 {
		return nil, errors.ErrIDInvalid
	}

	if userID <= 0 {
		return nil, errors.ErrUserIDRequired
	}

	book, err := bc.BookRepository.FindOne(id)
	if err != nil {
		return nil, err
	}

	if book.UserID != userID {
		return nil, errors.ErrUserMismatch
	}

	return book, nil
}

func (bc *BookControl) FetchMany(userID uint) ([]FullBook, error) {
	if userID <= 0 {
		return nil, errors.ErrUserIDRequired
	}

	return bc.BookRepository.FindMany(userID)
}
