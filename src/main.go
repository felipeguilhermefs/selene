package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/felipeguilhermefs/selene/views"
)

type server struct {
	router *mux.Router
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func newServer() (*server, error) {
	router := mux.NewRouter()

	view, err := views.NewView("books")
	if err != nil {
		return nil, err
	}
	router.Handle("/", view)

	return &server{
		router: router,
	}, nil
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
