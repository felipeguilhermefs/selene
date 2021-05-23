package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/felipeguilhermefs/selene/handlers"
	"github.com/felipeguilhermefs/selene/middlewares"
	"github.com/felipeguilhermefs/selene/services"
	"github.com/felipeguilhermefs/selene/view"
)

func NewRouter(mdw *middlewares.Middlewares, srvcs *services.Services) http.Handler {
	router := mux.NewRouter()

	registerRoutes(router, mdw.Login, srvcs)

	return mdw.CSRF(router.ServeHTTP)
}

func registerRoutes(
	router *mux.Router,
	loginMdw middlewares.Middleware,
	srvcs *services.Services,
) {
	signupView := view.NewView("signup")
	signupPage := handlers.HandleSignupPage(signupView)
	router.HandleFunc("/signup", signupPage).Methods("GET")
	signup := handlers.HandleSignup(signupView, srvcs.Auth)
	router.HandleFunc("/signup", signup).Methods("POST")

	loginView := view.NewView("login")
	loginPage := handlers.HandleLoginPage(loginView)
	router.HandleFunc("/login", loginPage).Methods("GET")
	login := handlers.HandleLogin(loginView, srvcs.Auth)
	router.HandleFunc("/login", login).Methods("POST")

	logout := handlers.HandleLogout(srvcs.Auth)
	router.HandleFunc("/logout", loginMdw(logout)).Methods("POST")

	booksView := view.NewView("books")
	booksPage := handlers.HandleBooksPage(booksView, srvcs.Book)
	router.HandleFunc("/books", loginMdw(booksPage)).Methods("GET")

	newBookView := view.NewView("new_book")
	newBookPage := handlers.HandleNewBookPage(newBookView)
	router.HandleFunc("/books/new", loginMdw(newBookPage)).Methods("GET")
	newBook := handlers.HandleNewBook(newBookView, srvcs.Book)
	router.HandleFunc("/books/new", loginMdw(newBook)).Methods("POST")

	bookView := view.NewView("book")
	editBookPage := handlers.HandleBookPage(bookView, srvcs.Book)
	router.HandleFunc("/books/{id:[0-9]+}", loginMdw(editBookPage)).Methods("GET")
	editBook := handlers.HandleEditBook(bookView, srvcs.Book)
	router.HandleFunc("/books/{id:[0-9]+}/edit", loginMdw(editBook)).Methods("POST")
	deleteBook := handlers.HandleDeleteBook(bookView, srvcs.Book)
	router.HandleFunc("/books/{id:[0-9]+}/delete", loginMdw(deleteBook)).Methods("POST")

	notFoundView := view.NewView("404")
	router.NotFoundHandler = handlers.HandleNotFound(notFoundView)

	errorView := view.NewView("error")
	router.MethodNotAllowedHandler = handlers.HandleError(
		errorView,
		http.StatusMethodNotAllowed,
		"Sorry, this HTTP method is not allowed in this route.",
	)
}
