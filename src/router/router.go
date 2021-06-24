package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/felipeguilhermefs/selene/handlers"
	"github.com/felipeguilhermefs/selene/middlewares"
)

type Router = http.Handler

func New(
	mdw *middlewares.Middlewares,
	hdlrs *handlers.Handlers,
) Router {
	router := mux.NewRouter()

	registerRoutes(router, mdw.Login, hdlrs)

	return mdw.CSRF(mdw.SecHeaders(router.ServeHTTP))
}

func registerRoutes(
	router *mux.Router,
	loginMdw middlewares.Middleware,
	hdlrs *handlers.Handlers,
) {
	router.HandleFunc("/signup", hdlrs.SignupPage).Methods("GET")
	router.HandleFunc("/signup", hdlrs.Signup).Methods("POST")
	router.HandleFunc("/login", hdlrs.LoginPage).Methods("GET")
	router.HandleFunc("/login", hdlrs.Login).Methods("POST")
	router.HandleFunc("/logout", loginMdw(hdlrs.Logout)).Methods("POST")
	router.HandleFunc("/books", loginMdw(hdlrs.BooksPage)).Methods("GET")
	router.HandleFunc("/books/new", loginMdw(hdlrs.NewBookPage)).Methods("GET")
	router.HandleFunc("/books/new", loginMdw(hdlrs.NewBook)).Methods("POST")
	router.HandleFunc("/books/{id:[0-9]+}", loginMdw(hdlrs.BookPage)).Methods("GET")
	router.HandleFunc("/books/{id:[0-9]+}/edit", loginMdw(hdlrs.EditBook)).Methods("POST")
	router.HandleFunc("/books/{id:[0-9]+}/delete", loginMdw(hdlrs.DeleteBook)).Methods("POST")

	router.NotFoundHandler = hdlrs.NotFound

	router.MethodNotAllowedHandler = hdlrs.Error
}
