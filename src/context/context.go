package context

import (
	"context"
	"net/http"

	"github.com/felipeguilhermefs/selene/core/auth"
)

const (
	userKey privateKey = "user"
)

type privateKey string

// WithUser Assign user to context
func WithUser(r *http.Request, user *auth.User) context.Context {
	return context.WithValue(r.Context(), userKey, user)
}

// User retrieve user from context
func User(r *http.Request) *auth.User {
	userCtx := r.Context().Value(userKey)
	if userCtx == nil {
		return nil
	}

	if user, ok := userCtx.(*auth.User); ok {
		return user
	}

	return nil
}
