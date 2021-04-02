package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/felipeguilhermefs/selene/infra/errors"
)

func run() error {
	cfg, err := LoadConfig()
	if err != nil {
		return errors.Wrap(err, "Loading config")
	}

	server, err := NewServer(cfg)
	if err != nil {
		return errors.Wrap(err, "Setting up Server")
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
