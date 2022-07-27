package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.CleanPath)
	router.Use(middleware.StripSlashes)
	router.Use(middleware.Recoverer)

	for _, md := range middlewares {
		router.Use(md)
	}

	for _, r := range routes {
		router.Method(r.Method, r.Path, r.Handler)
	}

	router.NotFound(notFound.ServeHTTP)

	return router
}
