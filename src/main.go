package main

import (
	"fmt"
	"log"
	"net/http"
)

func run() error {
	cfg, err := LoadConfig()
	if err != nil {
		return WrapError(err, "Loading config")
	}

	server, err := NewServer()
	if err != nil {
		return WrapError(err, "Setting up Server")
	}

	log.Printf("Server started at :%d...\n", cfg.Port)

	return http.ListenAndServe(
		fmt.Sprintf(":%d", cfg.Port),
		server,
	)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
