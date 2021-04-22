package services

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/felipeguilhermefs/selene/infra/config"
	"github.com/felipeguilhermefs/selene/infra/errors"
)

// SecretService handles operations over a secret key
type SecretService interface {
	Generate(password string) (string, error)
	Compare(secret, password string) error
}

// newSecretService creates a new instance of SecretService
func newSecretService(cfg *config.Config) SecretService {
	return &secretService{
		pwPepper: cfg.Sec.Pepper,
	}
}

type secretService struct {
	pwPepper string
}

// Generate creates a salt and pepper secret
func (ss *secretService) Generate(password string) (string, error) {
	raw := []byte(password + ss.pwPepper)
	hashed, err := bcrypt.GenerateFromPassword(raw, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// Compare compare a salt and pepper secret
func (ss *secretService) Compare(secret, password string) error {
	secretBytes := []byte(secret)
	passwordBytes := []byte(password + ss.pwPepper)

	err :=  bcrypt.CompareHashAndPassword(secretBytes, passwordBytes)

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return errors.ErrCredentialsInvalid
	}

	return err
}
