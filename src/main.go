package main

import (
	"log"

	"github.com/felipeguilhermefs/selene/infra/config"
	"github.com/felipeguilhermefs/selene/infra/errors"
)

func run() error {
	cfg, err := config.LoadFromFile("config.json")
	if err != nil {
		return errors.Wrap(err, "Loading config")
	}

	server, err := NewServer(cfg)
	if err != nil {
		return errors.Wrap(err, "Setting up Server")
	}

	server.RegisterRoutes()

	return server.Start()
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
