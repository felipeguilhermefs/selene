package server

import "time"

type ServerConfig interface {
	Port() int
	// ReadTimeout timeout for reading an entire request, including body
	ReadTimeout() time.Duration
	// WriteTimeout timeout for writing response
	WriteTimeout() time.Duration
	// IdleTimeout connection timeout receiving keep alives|
	IdleTimeout() time.Duration
}
