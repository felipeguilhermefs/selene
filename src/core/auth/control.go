package auth

type AuthControl struct {
	UserRepository  UserRepository
	EmailNormalizer EmailNormalizer
	PasswordControl PasswordControl
}

func (uc *AuthControl) Add(user *NewUser) error {
	if user.Email == "" {
		return ErrEmailRequired
	}

	normalizedEmail, err := uc.EmailNormalizer.Normalize(user.Email)
	if err != nil {
		return err
	}

	encriptedPassword, err := uc.PasswordControl.Generate(user.Password)
	if err != nil {
		return err
	}

	user.Email = normalizedEmail
	user.Password = encriptedPassword

	return uc.UserRepository.Add(user)
}

func (uc *AuthControl) FetchOne(email string) (*User, error) {
	fullUser, err := uc.findOne(email)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:    fullUser.ID,
		Name:  fullUser.Name,
		Email: fullUser.Email,
	}, nil
}

func (uc *AuthControl) Verify(email, password string) (*User, error) {
	fullUser, err := uc.findOne(email)
	if err != nil {
		return nil, err
	}

	err = uc.PasswordControl.Compare(fullUser.Password, password)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:    fullUser.ID,
		Name:  fullUser.Name,
		Email: fullUser.Email,
	}, nil
}

func (uc *AuthControl) findOne(email string) (*FullUser, error) {
	normalizedEmail, err := uc.EmailNormalizer.Normalize(email)
	if err != nil {
		return nil, err
	}

	return uc.UserRepository.FindOne(normalizedEmail)
}
