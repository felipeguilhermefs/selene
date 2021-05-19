package errors

const (
	// ErrNoCSRFField no CSRF field implemented
	ErrNoCSRFField PrivateError = "No CSRF field implemented"

	// ErrNotLoggedIn user not logged in
	ErrNotLoggedIn PrivateError = "Not logged in"

	// ErrUserIDRequired user id required
	ErrUserIDRequired PrivateError = "UserID is required"

	// ErrIDInvalid book id required
	ErrIDInvalid PrivateError = "ID is invalid"
)

type PrivateError string

func (e PrivateError) Error() string {
	return string(e)
}
