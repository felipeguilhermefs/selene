package services

import (
	"regexp"
	"strings"

	"github.com/felipeguilhermefs/selene/infra/errors"
	"github.com/felipeguilhermefs/selene/models"
	"github.com/felipeguilhermefs/selene/repositories"
)

// newUserService creates a new instance of UserService
func newUserService(
	ur repositories.UserRepository,
	ss SecretService,
) *userService {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,16}$`)

	return &userService{
		repository: ur,
		secretSrvc: ss,
		emailRegex: emailRegex,
	}
}

type userService struct {
	repository repositories.UserRepository
	secretSrvc SecretService
	emailRegex *regexp.Regexp
}

func (us *userService) Create(user *models.User) error {
	if user.Email == "" {
		return errors.ErrEmailRequired
	}

	if !us.emailRegex.MatchString(user.Email) {
		return errors.ErrEmailInvalid
	}

	if len(user.Password) < 8 {
		return errors.ErrPasswordTooShort
	}

	hashed, err := us.secretSrvc.Generate(user.Password)
	if err != nil {
		return err
	}

	user.Email = us.normalizeEmail(user.Email)
	user.Secret = hashed
	user.Password = ""

	return us.repository.Create(user)
}

func (us *userService) normalizeEmail(input string) string {
	return strings.ToLower(strings.TrimSpace(input))
}

func (us *userService) byEmail(email string) (*models.User, error) {
	normalized := us.normalizeEmail(email)

	if !us.emailRegex.MatchString(normalized) {
		return nil, errors.ErrEmailInvalid
	}

	return us.repository.ByEmail(normalized)
}
