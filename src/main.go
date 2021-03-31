package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	// "github.com/felipeguilhermefs/selene/views"
)

type server struct {
	router *mux.Router
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) handleBooksPage() http.HandlerFunc {
	type book struct {
		ID     uint
		Title  string
		Author string
		Tags   string
	}

	books := struct {
		Yield []book
	}{
		Yield: []book{
			{1, "The Hobbit", "JRR Tolkien", "adventure, fantasy"},
			{2, "Do Androids Dream of Electric Sheep?", "Philip K. Dick", "sci-fi, philosophical"},
			{3, "1984", "George Orwell", "dystopian, political fiction"},
		},
	}

	return HandleTemplate("books", func(r *http.Request) (interface{}, error) {
		return books, nil
	})
}

func newServer() (*server, error) {
	router := mux.NewRouter()

	s := server{
		router: router,
	}

	router.Handle("/", s.handleBooksPage())

	return &s, nil
}

func run() error {
	server, err := newServer()
	if err != nil {
		return err
	}

	port := 8000

	log.Printf("Server started at :%d...\n", port)

	return http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		server,
	)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
