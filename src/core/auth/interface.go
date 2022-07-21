package auth

import "github.com/felipeguilhermefs/selene/boundary/postgres"

type UserRepository interface {
	Create(user *postgres.User) error
	ByEmail(email string) (*postgres.User, error)
}
