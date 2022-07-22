package auth

type AuthControl struct {
	UserRepository    UserRepository
	EmailNormalizer   EmailNormalizer
	PasswordEncripter PasswordEncripter
}

func (uc *AuthControl) Add(user *NewUser) error {
	if user.Email == "" {
		return ErrEmailRequired
	}

	normalizedEmail, err := uc.EmailNormalizer.Normalize(user.Email)
	if err != nil {
		return err
	}

	encriptedPassword, err := uc.PasswordEncripter.Encript(user.Password)
	if err != nil {
		return err
	}

	user.Email = normalizedEmail
	user.Password = encriptedPassword

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
