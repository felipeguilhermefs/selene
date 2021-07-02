package auth

import "net/http"

// User represents an user of the system
type User struct {
	ID uint
}

// Request wraps an http.Request with the authenticated User
type Request struct {
	Request *http.Request
	User    *User
}

// Handler is analogous to http.HandlerFunc, but accepts the auth.Request
type Handler = func(w http.ResponseWriter, r *Request)

// Middleware simple http.Handler to http.Handler middleware
type Middleware = func(next Handler) http.Handler

// UserGetter function to figure out user from http.Request
type UserGetter = func(r *http.Request) (*User, error)

// ErrorHandler callback when user is not found
type ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error)

func NewMiddleware(
	getUser UserGetter,
	errorHandler ErrorHandler,
) Middleware {

	return func(next Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			user, err := getUser(r)
			if err != nil {
				errorHandler(w, r, err)
				return
			}

			next(w, &Request{
				Request: r,
				User:    user,
			})
		})
	}
}
