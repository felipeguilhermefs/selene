package context

import (
	"context"
	"net/http"

	"github.com/felipeguilhermefs/selene/boundary"
)

const (
	userKey privateKey = "user"
)

type privateKey string

// WithUser Assign user to context
func WithUser(r *http.Request, user *boundary.User) context.Context {
	return context.WithValue(r.Context(), userKey, user)
}

// User retrieve user from context
func User(r *http.Request) *boundary.User {
	userCtx := r.Context().Value(userKey)
	if userCtx == nil {
		return nil
	}

	if user, ok := userCtx.(*boundary.User); ok {
		return user
	}

	return nil
}
