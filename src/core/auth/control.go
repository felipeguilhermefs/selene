package auth

import (
	"regexp"
	"strings"

	"github.com/felipeguilhermefs/selene/infra/errors"
)

type AuthControl struct {
	UserRepository UserRepository
	EmailRegex     *regexp.Regexp
}

func (uc *AuthControl) Add(user *NewUser) error {
	if strings.TrimSpace(user.Password) == "" {
		return errors.ErrPasswordTooShort
	}

	normalizedEmail := uc.normalizeEmail(user.Email)

	if !uc.EmailRegex.MatchString(normalizedEmail) {
		return errors.ErrEmailInvalid
	}

	user.Email = normalizedEmail

	return uc.UserRepository.Add(user)
}

func (uc *AuthControl) FetchOne(email string) (*FullUser, error) {
	normalized := uc.normalizeEmail(email)

	if !uc.EmailRegex.MatchString(normalized) {
		return nil, errors.ErrEmailInvalid
	}

	return uc.UserRepository.FindOne(normalized)
}

func (uc *AuthControl) normalizeEmail(input string) string {
	return strings.ToLower(strings.TrimSpace(input))
}
