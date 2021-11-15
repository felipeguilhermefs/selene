package services

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/felipeguilhermefs/selene/infra/errors"
	"github.com/felipeguilhermefs/selene/infrastructure/config"
)

// PasswordService handles operations over a secret key
type PasswordService interface {
	Generate(password string) (string, error)
	Compare(secret, password string) error
}

// newPasswordService creates a new instance of PasswordService
func newPasswordService(cfg config.ConfigStore) PasswordService {
	return &passwordService{
		minLen: cfg.GetInt("SELENE_PW_MIN_LEN", 8),
		pepper: cfg.GetSecret("SELENE_PW_PEPPER", "PepperWith64Chars..............................................."),
	}
}

type passwordService struct {
	minLen int
	pepper string
}

// Generate creates a salt and pepper password
func (ss *passwordService) Generate(password string) (string, error) {
	if len(password) < ss.minLen {
		return "", errors.ErrPasswordTooShort
	}

	raw := []byte(password + ss.pepper)
	hashed, err := bcrypt.GenerateFromPassword(raw, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// Compare compare a salt and pepper password
func (ss *passwordService) Compare(secret, password string) error {
	secretBytes := []byte(secret)
	passwordBytes := []byte(password + ss.pepper)

	err := bcrypt.CompareHashAndPassword(secretBytes, passwordBytes)

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return errors.ErrCredentialsInvalid
	}

	return err
}
