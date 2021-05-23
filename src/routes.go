package main

import (
	"github.com/felipeguilhermefs/selene/handlers"
	"github.com/felipeguilhermefs/selene/view"
)

// RegisterRoutes register all server routes
func (s *Server) RegisterRoutes() {
	signupView := view.NewView("signup")
	signupPage := handlers.HandleSignupPage(signupView)
	s.router.HandleFunc("/signup", signupPage).Methods("GET")
	signup := handlers.HandleSignup(signupView, s.services.Auth)
	s.router.HandleFunc("/signup", signup).Methods("POST")

	loginView := view.NewView("login")
	loginPage := handlers.HandleLoginPage(loginView)
	s.router.HandleFunc("/login", loginPage).Methods("GET")
	login := handlers.HandleLogin(loginView, s.services.Auth)
	s.router.HandleFunc("/login", login).Methods("POST")

	logout := handlers.HandleLogout(s.services.Auth)
	s.router.HandleFunc("/logout", s.middlewares.Login(logout)).Methods("POST")

	booksView := view.NewView("books")
	booksPage := handlers.HandleBooksPage(booksView, s.services.Book)
	s.router.HandleFunc("/books", s.middlewares.Login(booksPage)).Methods("GET")

	newBookView := view.NewView("new_book")
	newBookPage := handlers.HandleNewBookPage(newBookView)
	s.router.HandleFunc("/books/new", s.middlewares.Login(newBookPage)).Methods("GET")
	newBook := handlers.HandleNewBook(newBookView, s.services.Book)
	s.router.HandleFunc("/books/new", s.middlewares.Login(newBook)).Methods("POST")

	bookView := view.NewView("book")
	editBookPage := handlers.HandleBookPage(bookView, s.services.Book)
	s.router.HandleFunc("/books/{id:[0-9]+}", s.middlewares.Login(editBookPage)).Methods("GET")
	editBook := handlers.HandleEditBook(bookView, s.services.Book)
	s.router.HandleFunc("/books/{id:[0-9]+}/edit", s.middlewares.Login(editBook)).Methods("POST")
	deleteBook := handlers.HandleDeleteBook(bookView, s.services.Book)
	s.router.HandleFunc("/books/{id:[0-9]+}/delete", s.middlewares.Login(deleteBook)).Methods("POST")

	notFoundView := view.NewView("404")
	s.router.NotFoundHandler = handlers.HandleNotFound(notFoundView)

	methodNotAllowedView := view.NewView("405")
	s.router.MethodNotAllowedHandler = handlers.HandleMethodNotAllowed(methodNotAllowedView)
}
