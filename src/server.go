package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"

	"github.com/felipeguilhermefs/selene/infra/config"
	"github.com/felipeguilhermefs/selene/infra/database"
	"github.com/felipeguilhermefs/selene/infra/errors"
	"github.com/felipeguilhermefs/selene/infra/session"
	"github.com/felipeguilhermefs/selene/repositories"
	"github.com/felipeguilhermefs/selene/services"
)

// Server represents all insfrastructure used in this server app
type Server struct {
	repositories *repositories.Repositories
	services     *services.Services
	server       *http.Server
	router       *mux.Router
}

// Start start listening and serving requests
func (s *Server) Start() error {
	log.Printf("Server started at %v...\n", s.server.Addr)
	return s.server.ListenAndServe()
}

// NewServer creates a new server instance
func NewServer(cfg *config.Config) (*Server, error) {
	router := mux.NewRouter()

	db, err := database.ConnectPostgres(&cfg.DB)
	if err != nil {
		return nil, errors.Wrap(err, "Connecting to Postgres")
	}

	sessionStore := session.NewCookieStore(&cfg.Sec.Session)

	repos := repositories.NewRepositories(db, sessionStore)
	if err := repos.AutoMigrate(); err != nil {
		return nil, errors.Wrap(err, "Migrating repositories")
	}

	srvcs := services.NewServices(cfg, repos)

	CSRF := csrf.Protect(
		[]byte(cfg.Sec.CSRF),
		csrf.SameSite(csrf.SameSiteStrictMode),
	)

	return &Server{
		repositories: repos,
		router:       router,
		services:     srvcs,
		server: &http.Server{
			Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
			ReadTimeout:  cfg.Server.ReadTimeout(),
			WriteTimeout: cfg.Server.WriteTimeout(),
			IdleTimeout:  cfg.Server.IdleTimeout(),
			Handler:      CSRF(router),
		},
	}, nil
}
