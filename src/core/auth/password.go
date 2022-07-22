package auth

import "golang.org/x/crypto/bcrypt"

type PasswordEncripter interface {
	Encript(password string) (string, error)
}

type PasswordComparer interface {
	Compare(secret, password string) error
}

type PasswordControl struct {
	MinLen int
	Pepper string
}

func (pc *PasswordControl) Encript(password string) (string, error) {
	if len(password) < pc.MinLen {
		return "", ErrPasswordTooShort
	}

	raw := []byte(password + pc.Pepper)
	hashed, err := bcrypt.GenerateFromPassword(raw, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func (pc *PasswordControl) Compare(secret, password string) error {
	secretBytes := []byte(secret)
	passwordBytes := []byte(password + pc.Pepper)

	err := bcrypt.CompareHashAndPassword(secretBytes, passwordBytes)

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return ErrCredentialsInvalid
	}

	return err
}
