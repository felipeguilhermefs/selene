package main

import (
	"log"

	"github.com/felipeguilhermefs/selene/infra/config"
	"github.com/felipeguilhermefs/selene/server"
)

func run() error {
	cfg, err := config.LoadFromFile("config.json")
	if err != nil {
		return err
	}

	s, err := server.New(cfg)
	if err != nil {
		return err
	}

	return s.Start()
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
