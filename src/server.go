package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Server represents all insfrastructure used in this server app
type Server struct {
	router *mux.Router
}

// ServeHTTP just delegates to router so we can start Server in http.ListenAndServe
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) handleBooksPage() http.HandlerFunc {
	type book struct {
		ID     uint
		Title  string
		Author string
		Tags   string
	}

	data := struct {
		Yield []book
	}{
		Yield: []book{
			{1, "The Hobbit", "JRR Tolkien", "adventure, fantasy"},
			{2, "Do Androids Dream of Electric Sheep?", "Philip K. Dick", "sci-fi, philosophical"},
			{3, "1984", "George Orwell", "dystopian, political fiction"},
		},
	}

	return HandleTemplate("books", func(r *http.Request) (interface{}, error) {
		return data, nil
	})
}

func (s *Server) handleBookPage() http.HandlerFunc {
	type book struct {
		ID       uint
		Title    string
		Author   string
		Comments string
		Tags     string
	}

	data := struct {
		Yield book
	}{
		Yield: book{
			ID:       2,
			Title:    "Do Androids Dream of Electric Sheep?",
			Author:   "Philip K. Dick",
			Comments: "quero muito ler esse",
			Tags:     "sci-fi, philosophical",
		},
	}

	return HandleTemplate("book", func(r *http.Request) (interface{}, error) {
		return data, nil
	})
}

func (s *Server) handleNewBookPage() http.HandlerFunc {
	type book struct {
		Title    string
		Author   string
		Comments string
		Tags     string
	}

	data := struct {
		Yield book
	}{
		Yield: book{
			Title:    "American Gods",
			Author:   "Neil Gaiman",
			Comments: "opa",
			Tags:     "fantasy",
		},
	}

	return HandleTemplate("new_book", func(r *http.Request) (interface{}, error) {
		return data, nil
	})
}

// NewServer creates a new server instance
func NewServer() (*Server, error) {
	router := mux.NewRouter()

	s := Server{
		router: router,
	}

	router.HandleFunc("/books", s.handleBooksPage()).Methods("GET")
	router.HandleFunc("/books/{id:[0-9]+}", s.handleBookPage()).Methods("GET")
	router.HandleFunc("/books/new", s.handleNewBookPage()).Methods("GET")

	return &s, nil
}