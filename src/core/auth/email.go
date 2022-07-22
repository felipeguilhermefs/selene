package auth

import (
	"regexp"
	"strings"
)

var emailRegex *regexp.Regexp = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,16}$`)

type EmailNormalizer struct {
}

func (en *EmailNormalizer) Normalize(email string) (string, error) {
	normalizedEmail := strings.ToLower(strings.TrimSpace(email))

	if !emailRegex.MatchString(normalizedEmail) {
		return "", ErrEmailInvalid
	}

	return normalizedEmail, nil
}
