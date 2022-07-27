package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Route struct {
	Method  string
	Path    string
	Handler http.Handler
}

type Middleware func(http.Handler) http.Handler

func New(
	routes []Route,
	middlewares []Middleware,
	notFound http.Handler,
) http.Handler {
	// r.Use(middleware.Logger)
	router := chi.NewRouter()

	for _, r := range routes {
		router.Method(r.Method, r.Path, r.Handler)
	}

	for _, md := range middlewares {
		router.Use(md)
	}

	router.NotFound(notFound.ServeHTTP)

	return router
}
