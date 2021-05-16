package main

import (
	"github.com/felipeguilhermefs/selene/handlers"
	"github.com/felipeguilhermefs/selene/view"
)

// RegisterRoutes register all server routes
func (s *Server) RegisterRoutes() {
	loginView := view.NewView("login")
	signupView := view.NewView("signup")
	booksView := view.NewView("books")
	newBookView := view.NewView("new_book")
	bookView := view.NewView("book")

	signupPage := handlers.HandleSignupPage(signupView)
	s.router.HandleFunc("/signup", signupPage).Methods("GET")

	signup := handlers.HandleSignup(signupView, s.services.Auth)
	s.router.HandleFunc("/signup", signup).Methods("POST")

	loginPage := handlers.HandleLoginPage(loginView)
	s.router.HandleFunc("/login", loginPage).Methods("GET")

	login := handlers.HandleLogin(loginView, s.services.Auth)
	s.router.HandleFunc("/login", login).Methods("POST")

	logout := handlers.HandleLogout(s.services.Auth)
	s.router.HandleFunc("/logout", logout).Methods("POST")

	booksPage := handlers.HandleBooksPage(booksView, s.services.Auth, s.services.Book)
	s.router.HandleFunc("/books", booksPage).Methods("GET")

	newBookPage := handlers.HandleNewBookPage(newBookView, s.services.Auth)
	s.router.HandleFunc("/books/new", newBookPage).Methods("GET")

	newBook := handlers.HandleNewBook(newBookView, s.services.Auth, s.services.Book)
	s.router.HandleFunc("/books/new", newBook).Methods("POST")

	editBookPage := handlers.HandleBookPage(bookView, s.services.Auth, s.services.Book)
	s.router.HandleFunc("/books/{id:[0-9]+}", editBookPage).Methods("GET")

	editBook := handlers.HandleEditBook(bookView, s.services.Auth, s.services.Book)
	s.router.HandleFunc("/books/{id:[0-9]+}/edit", editBook).Methods("POST")

	deleteBook := handlers.HandleDeleteBook(bookView, s.services.Auth, s.services.Book)
	s.router.HandleFunc("/books/{id:[0-9]+}/delete", deleteBook).Methods("POST")
}
