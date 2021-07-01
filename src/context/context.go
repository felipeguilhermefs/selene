package context

import (
	"context"
	"net/http"

	"github.com/felipeguilhermefs/selene/models"
)

const (
	UserKey privateKey = "user"
)

type privateKey string

// WithUser Assign user to context
func WithUser(r *http.Request, user *models.User) context.Context {
	return context.WithValue(r.Context(), UserKey, user)
}
