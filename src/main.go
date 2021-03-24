package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/felipeguilhermefs/selene/views"
)

func newServer() http.Handler {
	router := mux.NewRouter()

	view, err := views.NewView("home")
	if err != nil {
		panic(err)
	}
	router.Handle("/", view)

	return router
}

func run() error {
	server := newServer()
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
