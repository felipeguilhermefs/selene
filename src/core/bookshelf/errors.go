package bookshelf

const (
	ErrIDInvalid      BookshelfError = "ID is invalid"
	ErrUserIDRequired BookshelfError = "UserID is required"
	ErrTitleRequired  BookshelfError = "Book title is required"
	ErrAuthorRequired BookshelfError = "Book author is required"
	ErrUserMismatch   BookshelfError = "User mismatch, should be same"
	ErrBookNotFound   BookshelfError = "Book not found"
)

type BookshelfError string

func (e BookshelfError) Error() string {
	return string(e)
}
