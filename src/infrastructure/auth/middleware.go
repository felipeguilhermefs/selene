package auth

import (
	"context"
	"net/http"
)

const (
	UserKey privateKey = "user"
)

type privateKey string

func NewMiddleware(authService AuthService) func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			user, err := authService.GetUser(r)
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}

			ctx := context.WithValue(r.Context(), UserKey, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
