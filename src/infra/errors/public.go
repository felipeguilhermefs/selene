package errors

const (
	// ErrUserNotFound user not found
	ErrUserNotFound PublicError = "User not found"

	// ErrBookNotFound book not found
	ErrBookNotFound PublicError = "Book not found"

	// ErrEmailRequired Email is required
	ErrEmailRequired PublicError = "Email is required"

	// ErrEmailInvalid Email is invalid
	ErrEmailInvalid PublicError = "Email is invalid"

	// ErrPasswordTooShort Password too short
	ErrPasswordTooShort PublicError = "Password too short"

	// ErrCredentialsInvalid username or password not valid
	ErrCredentialsInvalid PublicError = "Invalid credentials"

	// ErrTitleRequired book title required
	ErrTitleRequired PublicError = "Book title is required"

	// ErrAuthorRequired book author required
	ErrAuthorRequired PublicError = "Book author is required"
)

type PublicError string

func (e PublicError) Error() string {
	return string(e)
}
