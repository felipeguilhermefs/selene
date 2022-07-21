package auth

type NewUser struct {
	Name     string
	Email    string
	Password string
}

type UserAdder interface {
	Add(user *NewUser) error
}

type User struct {
	ID    uint
	Name  string
	Email string
}

type UserFetcher interface {
	FetchOne(email string) (*User, error)
}

type FullUser struct {
	ID       uint
	Name     string
	Email    string
	Password string
}

type UserRepository interface {
	Add(user *NewUser) error
	FindOne(email string) (*FullUser, error)
}
