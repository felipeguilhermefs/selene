package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func newServer() http.Handler {
	router := mux.NewRouter()

	view, err := NewView("home")
	if err != nil {
		panic(err)
	}
    router.Handle("/", view)

	return router
}

func run() error {
	server := newServer()
	port := 8000;

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
