package errors

const (
	// ErrUserNotFound user not found
	ErrUserNotFound publicError = "User not found"

	// ErrBookNotFound book not found
	ErrBookNotFound publicError = "Book not found"

	// ErrEmailRequired Email is required
	ErrEmailRequired publicError = "Email is required"

	// ErrEmailInvalid Email is invalid
	ErrEmailInvalid publicError = "Email is invalid"

	// ErrPasswordTooShort Password too short
	ErrPasswordTooShort publicError = "Password too short"

	// ErrCredentialsInvalid username or password not valid
	ErrCredentialsInvalid publicError = "Invalid credentials"

	// ErrTitleRequired book title required
	ErrTitleRequired publicError = "Book title is required"

	// ErrAuthorRequired book author required
	ErrAuthorRequired publicError = "Book author is required"
)

type publicError string

func (e publicError) Error() string {
	return string(e)
}
