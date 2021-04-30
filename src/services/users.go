package services

import (
	"regexp"
	"strings"

	"github.com/felipeguilhermefs/selene/infra/errors"
	"github.com/felipeguilhermefs/selene/models"
	"github.com/felipeguilhermefs/selene/repositories"
)

// UserService handle operations over users
type UserService interface {
	Create(user *models.User) error
	Authenticate(email, password string) (*models.User, error)
}

// newUserService creates a new instance of UserService
func newUserService(
	ur repositories.UserRepository,
	ss SecretService,
) UserService {
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

func (us *userService) Authenticate(email, password string) (*models.User, error) {
	user, err := us.byEmail(email)
	if err != nil {
		return nil, err
	}

	err = us.secretSrvc.Compare(user.Secret, password)
	if err != nil {
		return nil, err
	}

	return user, nil
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
