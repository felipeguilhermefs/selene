package middlewares

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/context"
	"github.com/felipeguilhermefs/selene/services"
)

func newLoginMiddleware(authService services.AuthService) Middleware {

	return func(next http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {

			user, err := authService.GetUser(r)
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}

			ctx := context.WithUser(r, user)

			next(w, r.WithContext(ctx))
		}
	}
}
