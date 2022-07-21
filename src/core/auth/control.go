package auth

import (
	"strings"

	"github.com/felipeguilhermefs/selene/infra/errors"
)

type AuthControl struct {
	UserRepository  UserRepository
	EmailNormalizer EmailNormalizer
}

func (uc *AuthControl) Add(user *NewUser) error {
	if strings.TrimSpace(user.Password) == "" {
		return errors.ErrPasswordTooShort
	}

	normalizedEmail, err := uc.EmailNormalizer.Normalize(user.Email)
	if err != nil {
		return err
	}

	user.Email = normalizedEmail

	return uc.UserRepository.Add(user)
}

func (uc *AuthControl) FetchOne(email string) (*User, error) {
	normalizedEmail, err := uc.EmailNormalizer.Normalize(email)
	if err != nil {
		return nil, err
	}

	fullUser, err := uc.UserRepository.FindOne(normalizedEmail)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:    fullUser.ID,
		Name:  fullUser.Name,
		Email: fullUser.Email,
	}, nil
}
