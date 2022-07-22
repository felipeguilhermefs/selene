package services

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/felipeguilhermefs/selene/core/auth"
	"github.com/felipeguilhermefs/selene/infra/errors"
	"github.com/felipeguilhermefs/selene/infrastructure/config"
)

// PasswordService handles operations over a secret key
type PasswordService interface {
	Compare(secret, password string) error
}

// newPasswordService creates a new instance of PasswordService
func newPasswordService(cfg config.ConfigStore, passwordEncripter auth.PasswordEncripter) PasswordService {
	return &passwordService{
		passwordEncripter: passwordEncripter,
		pepper:            cfg.GetSecret("SELENE_PW_PEPPER", "PepperWith64Chars..............................................."),
	}
}

type passwordService struct {
	pepper            string
	passwordEncripter auth.PasswordEncripter
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
