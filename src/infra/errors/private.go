package errors

const (
	// ErrNoCSRFField no CSRF field implemented
	ErrNoCSRFField privateError = "No CSRF field implemented"

	// ErrNotLoggedIn user not logged in
	ErrNotLoggedIn privateError = "Not logged in"

	// ErrUserIDRequired user id required
	ErrUserIDRequired privateError = "UserID is required"

	// ErrIDInvalid book id required
	ErrIDInvalid privateError = "ID is invalid"
)

type privateError string

func (e privateError) Error() string {
	return string(e)
}
