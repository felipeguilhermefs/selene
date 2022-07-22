package auth

const (
	ErrPasswordTooShort      AuthError = "Password too short"
	ErrEmailRequired AuthError = "Email is required"
	ErrEmailInvalid AuthError = "Email is invalid"
	ErrUserNotFound AuthError = "User not found"
)

type AuthError string

func (e AuthError) Error() string {
	return string(e)
}
