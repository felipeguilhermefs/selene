package bookshelf

import "strings"

type BookControl struct {
	BookRepository BookRepository
}

func (bc *BookControl) Add(book *NewBook) error {
	if book.UserID <= 0 {
		return ErrUserIDRequired
	}

	if strings.TrimSpace(book.Title) == "" {
		return ErrTitleRequired
	}

	if strings.TrimSpace(book.Author) == "" {
		return ErrAuthorRequired
	}

	return bc.BookRepository.Insert(book)
}

func (bc *BookControl) Update(book *UpdatedBook) error {
	if book.ID <= 0 {
		return ErrIDInvalid
	}

	if book.UserID <= 0 {
		return ErrUserIDRequired
	}

	if strings.TrimSpace(book.Title) == "" {
		return ErrTitleRequired
	}

	if strings.TrimSpace(book.Author) == "" {
		return ErrAuthorRequired
	}

	current, err := bc.BookRepository.FindOne(book.ID)
	if err != nil {
		return err
	}

	if current.UserID != book.UserID {
		return ErrUserMismatch
	}

	return bc.BookRepository.Update(book)
}

func (bc *BookControl) Remove(userID uint, id uint) error {
	if id <= 0 {
		return ErrIDInvalid
	}

	if userID <= 0 {
		return ErrUserIDRequired
	}

	book, err := bc.BookRepository.FindOne(id)
	if err != nil {
		return err
	}

	if book.UserID != userID {
		return ErrUserMismatch
	}

	return bc.BookRepository.Delete(id)
}

func (bc *BookControl) FetchOne(userID, id uint) (*FullBook, error) {
	if id <= 0 {
		return nil, ErrIDInvalid
	}

	if userID <= 0 {
		return nil, ErrUserIDRequired
	}

	book, err := bc.BookRepository.FindOne(id)
	if err != nil {
		return nil, err
	}

	if book.UserID != userID {
		return nil, ErrUserMismatch
	}

	return book, nil
}

func (bc *BookControl) FetchMany(userID uint) ([]FullBook, error) {
	if userID <= 0 {
		return nil, ErrUserIDRequired
	}

	return bc.BookRepository.FindMany(userID)
}
