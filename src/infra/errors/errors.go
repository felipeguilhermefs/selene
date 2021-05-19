package errors

import "github.com/pkg/errors"

const (
	// errorSeparator is just a format suggar to separate each known error step
	errorSeparator = " >> "

	// ErrNoCSRFField no CSRF field implemented
	ErrNoCSRFField knownError = "No CSRF field implemented"

	// ErrNotLoggedIn user not logged in
	ErrNotLoggedIn knownError = "Not logged in"

	// ErrUserIDRequired user id required
	ErrUserIDRequired knownError = "UserID is required"

	// ErrIDInvalid book id required
	ErrIDInvalid knownError = "ID is invalid"
)

// Wrap Improve an error with a message giving it more context
func Wrap(err error, message string) error {
	return errors.Wrap(err, message+errorSeparator)
}

type knownError string

func (e knownError) Error() string {
	return string(e)
}
