package auth

import (
	"github.com/felipeguilhermefs/selene/infrastructure/config"
	"golang.org/x/crypto/bcrypt"
)

type PasswordControl interface {
	Generate(password string) (string, error)
	Compare(secret, password string) error
}

func NewPasswordControl(cfg config.ConfigStore) PasswordControl {
	return &bcriptPasswordControl{
		minLen: cfg.GetInt("SELENE_PW_MIN_LEN", 8),
		pepper: cfg.GetSecret("SELENE_PW_PEPPER", "PepperWith64Chars..............................................."),
	}
}

type bcriptPasswordControl struct {
	minLen int
	pepper string
}

func (bpc *bcriptPasswordControl) Generate(password string) (string, error) {
	if len(password) < bpc.minLen {
		return "", ErrPasswordTooShort
	}

	raw := []byte(password + bpc.pepper)
	hashed, err := bcrypt.GenerateFromPassword(raw, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func (bpc *bcriptPasswordControl) Compare(secret, password string) error {
	secretBytes := []byte(secret)
	passwordBytes := []byte(password + bpc.pepper)

	err := bcrypt.CompareHashAndPassword(secretBytes, passwordBytes)

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return ErrCredentialsInvalid
	}

	return err
}
