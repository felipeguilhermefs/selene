package errors

import "github.com/pkg/errors"

const (
	// errorSeparator is just a format suggar to separate each known error step
	errorSeparator = " >> "

	// ErrNoCSRFField no CSRF field implemented
	ErrNoCSRFField knownError = "No CSRF field implemented"

	// ErrNotFound resource not found
	ErrNotFound knownError = "Resource not found"

	// ErrEmailRequired Email is required
	ErrEmailRequired knownError = "Email is required"

	// ErrEmailInvalid Email is invalid
	ErrEmailInvalid knownError = "Email is invalid"

	// ErrPasswordTooShort Password too short
	ErrPasswordTooShort knownError = "Password too short"

	// ErrCredentialsInvalid username or password not valid
	ErrCredentialsInvalid knownError = "Invalid credentials"

	// ErrNotLoggedIn user not logged in
	ErrNotLoggedIn knownError = "Not logged in"

	// ErrUserIDRequired user id required
	ErrUserIDRequired knownError = "UserID is required"

	// ErrIDInvalid book id required
	ErrIDInvalid knownError = "ID is invalid"

	// ErrTitleRequired book title required
	ErrTitleRequired knownError = "Book title is required"

	// ErrAuthorRequired book author required
	ErrAuthorRequired knownError = "Book author is required"
)

// Wrap Improve an error with a message giving it more context
func Wrap(err error, message string) error {
	return errors.Wrap(err, message+errorSeparator)
}

type knownError string

func (e knownError) Error() string {
	return string(e)
}
