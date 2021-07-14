package server

import (
	"fmt"
	"net/http"

	"github.com/felipeguilhermefs/selene/infrastructure/config"
)

func New(cfg config.ConfigStore, router http.Handler) *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.GetInt("SELENE_SERVER_PORT", 8000)),
		ReadTimeout:  cfg.GetTime("SELENE_SERVER_READ_TIMEOUT", "15s"),
		WriteTimeout: cfg.GetTime("SELENE_SERVER_WRITE_TIMEOUT", "15s"),
		IdleTimeout:  cfg.GetTime("SELENE_SERVER_IDLE_TIMEOUT", "60s"),
		Handler:      router,
	}
}
