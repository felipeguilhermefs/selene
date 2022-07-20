package errors

const (
	// ErrNoCSRFField no CSRF field implemented
	ErrNoCSRFField PrivateError = "No CSRF field implemented"

	// ErrNotLoggedIn user not logged in
	ErrNotLoggedIn PrivateError = "Not logged in"
)

type PrivateError string

func (e PrivateError) Error() string {
	return string(e)
}
