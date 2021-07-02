package auth

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/models"
)

type Request struct {
	Request *http.Request
	User    *models.User
}

type Handler = func(w http.ResponseWriter, r *Request)

type Middleware = func(next Handler) http.Handler

type UserGetter = func(r *http.Request) (*models.User, error)

func NewMiddleware(getUser UserGetter) Middleware {

	return func(next Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			user, err := getUser(r)
			if err != nil {
				http.Redirect(w, r, "/login", http.StatusFound)
				return
			}

			next(w, &Request{
				Request: r,
				User:    user,
			})
		})
	}
}
