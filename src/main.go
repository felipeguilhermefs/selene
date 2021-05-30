package main

import (
	"log"

	"github.com/felipeguilhermefs/selene/infra/config"
)

func run() error {
	cfg, err := config.LoadFromFile("config.json")
	if err != nil {
		return err
	}

	server, err := NewServer(cfg)
	if err != nil {
		return err
	}

	return server.Start()
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
