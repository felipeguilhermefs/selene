package auth

import "github.com/felipeguilhermefs/selene/boundary"

type UserRepository interface {
	Create(user *boundary.User) error
	ByEmail(email string) (*boundary.User, error)
}
