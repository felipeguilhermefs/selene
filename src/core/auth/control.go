package auth

import (
	"regexp"
	"strings"

	"github.com/felipeguilhermefs/selene/infra/errors"
)

type UserControl struct {
	UserRepository UserRepository
	EmailRegex     *regexp.Regexp
}

func (uc *UserControl) Add(user *NewUser) error {
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

func (uc *UserControl) FetchOne(email string) (*FullUser, error) {
	normalized := uc.normalizeEmail(email)

	if !uc.EmailRegex.MatchString(normalized) {
		return nil, errors.ErrEmailInvalid
	}

	return uc.UserRepository.FindOne(normalized)
}

func (uc *UserControl) normalizeEmail(input string) string {
	return strings.ToLower(strings.TrimSpace(input))
}
