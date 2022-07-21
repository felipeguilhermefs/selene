package auth

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/context"
	"github.com/felipeguilhermefs/selene/core/auth"
	"github.com/felipeguilhermefs/selene/infrastructure/session"
)

func New(userFetcher auth.UserFetcher, sessionStore session.SessionStore) func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			email, err := sessionStore.GetUserID(r)
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}

			user, err := userFetcher.FetchOne(email)
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}

			ctx := context.WithUser(r, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
