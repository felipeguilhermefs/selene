package auth

import "golang.org/x/crypto/bcrypt"

type PasswordEncripter interface {
	Encript(password string) (string, error)
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
