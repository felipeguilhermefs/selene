package main

import (
	"errors"
	"log"
)

func run() error {
	return errors.New("coiso")
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
