package main

import (
	"fmt"
	"log"
	"net/http"
)

func run() error {
	server, err := NewServer()
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
