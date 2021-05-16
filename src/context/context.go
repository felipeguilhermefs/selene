package context

import (
	"context"
	"net/http"

	"github.com/felipeguilhermefs/selene/models"
)

const (
	userKey privateKey = "user"
)

type privateKey string

// WithUser Assign user to context
func WithUser(r *http.Request, user *models.User) context.Context {
	return context.WithValue(r.Context(), userKey, user)
}

// User retrieve user from context
func User(r *http.Request) *models.User {
	userCtx := r.Context().Value(userKey)
	if userCtx == nil {
		return nil
	}

	if user, ok := userCtx.(*models.User); ok {
		return user
	}

	return nil
}
