package auth

import "strings"

type AuthControl struct {
	UserRepository  UserRepository
	EmailNormalizer EmailNormalizer
}

func (uc *AuthControl) Add(user *NewUser) error {
	if user.Email == "" {
		return ErrEmailRequired
	}

	normalizedEmail, err := uc.EmailNormalizer.Normalize(user.Email)
	if err != nil {
		return err
	}

	if strings.TrimSpace(user.Password) == "" {
		return ErrPasswordTooShort
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
