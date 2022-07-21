package auth

type NewUser struct {
	Name     string
	Email    string
	Password string
}

type UserAdder interface {
	Add(user *NewUser) error
}

type FullUser struct {
	ID       uint
	Name     string
	Email    string
	Password string
}

type UserFetcher interface {
	FetchOne(email string) (*FullUser, error)
}

type UserRepository interface {
	Add(user *NewUser) error
	FindOne(email string) (*FullUser, error)
}
