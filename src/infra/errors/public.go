package errors

const (
	// ErrCredentialsInvalid username or password not valid
	ErrCredentialsInvalid PublicError = "Invalid credentials"
)

type PublicError string

func (e PublicError) Error() string {
	return string(e)
}
