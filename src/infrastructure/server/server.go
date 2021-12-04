package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/felipeguilhermefs/selene/infrastructure/config"
)

type Server interface {
	Serve() error
	Shutdown() error
}

type server struct {
	internal        http.Server
	shutdownTimeout time.Duration
}

func (srv *server) Serve() error {
	log.Printf("Server started at %v...\n", srv.internal.Addr)

	err := srv.internal.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (srv *server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), srv.shutdownTimeout)
	defer cancel()

	log.Printf("Server shutting down in %v", srv.shutdownTimeout)

	return srv.internal.Shutdown(ctx)
}

func New(cfg config.ConfigStore, router http.Handler) Server {
	return &server{
		internal: http.Server{
			Addr:         fmt.Sprintf(":%d", cfg.GetInt("SELENE_SERVER_PORT", 8000)),
			ReadTimeout:  cfg.GetTime("SELENE_SERVER_READ_TIMEOUT", "15s"),
			WriteTimeout: cfg.GetTime("SELENE_SERVER_WRITE_TIMEOUT", "15s"),
			IdleTimeout:  cfg.GetTime("SELENE_SERVER_IDLE_TIMEOUT", "60s"),
			Handler:      router,
		},
		shutdownTimeout: cfg.GetTime("SELENE_SERVER_SHUTDOWN_TIMEOUT", "10s"),
	}
}
