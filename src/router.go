package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/felipeguilhermefs/selene/handlers"
	"github.com/felipeguilhermefs/selene/middlewares"
	"github.com/felipeguilhermefs/selene/services"
	"github.com/felipeguilhermefs/selene/view"
)

func NewRouter(
	mdw *middlewares.Middlewares,
	srvcs *services.Services,
	views *view.Views,
) http.Handler {
	router := mux.NewRouter()

	registerRoutes(router, mdw.Login, srvcs, views)

	return mdw.CSRF(router.ServeHTTP)
}

func registerRoutes(
	router *mux.Router,
	loginMdw middlewares.Middleware,
	srvcs *services.Services,
	views *view.Views,
) {
	signupPage := handlers.HandleSignupPage(views.Signup)
	router.HandleFunc("/signup", signupPage).Methods("GET")

	signup := handlers.HandleSignup(views.Signup, srvcs.Auth)
	router.HandleFunc("/signup", signup).Methods("POST")

	loginPage := handlers.HandleLoginPage(views.Login)
	router.HandleFunc("/login", loginPage).Methods("GET")

	login := handlers.HandleLogin(views.Login, srvcs.Auth)
	router.HandleFunc("/login", login).Methods("POST")

	logout := handlers.HandleLogout(srvcs.Auth)
	router.HandleFunc("/logout", loginMdw(logout)).Methods("POST")

	booksPage := handlers.HandleBooksPage(views.Books, srvcs.Book)
	router.HandleFunc("/books", loginMdw(booksPage)).Methods("GET")

	newBookPage := handlers.HandleNewBookPage(views.NewBook)
	router.HandleFunc("/books/new", loginMdw(newBookPage)).Methods("GET")

	newBook := handlers.HandleNewBook(views.NewBook, srvcs.Book)
	router.HandleFunc("/books/new", loginMdw(newBook)).Methods("POST")

	editBookPage := handlers.HandleBookPage(views.EditBook, srvcs.Book)
	router.HandleFunc("/books/{id:[0-9]+}", loginMdw(editBookPage)).Methods("GET")

	editBook := handlers.HandleEditBook(views.EditBook, srvcs.Book)
	router.HandleFunc("/books/{id:[0-9]+}/edit", loginMdw(editBook)).Methods("POST")

	deleteBook := handlers.HandleDeleteBook(views.EditBook, srvcs.Book)
	router.HandleFunc("/books/{id:[0-9]+}/delete", loginMdw(deleteBook)).Methods("POST")

	router.NotFoundHandler = handlers.HandleNotFound(views.NotFound)

	router.MethodNotAllowedHandler = handlers.HandleError(
		views.Error,
		http.StatusMethodNotAllowed,
		"Sorry, this HTTP method is not allowed in this route.",
	)
}
