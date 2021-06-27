package router

import (
	"net/http"

	"github.com/gorilla/mux"
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
	router := mux.NewRouter()

	for _, r := range routes {
		router.Handle(r.Path, r.Handler).Methods(r.Method)
	}

	for _, md := range middlewares {
		router.Use(mux.MiddlewareFunc(md))
	}

	router.NotFoundHandler = notFound

	return router
}