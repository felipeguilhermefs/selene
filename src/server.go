package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	"github.com/felipeguilhermefs/selene/controllers"
	"github.com/felipeguilhermefs/selene/infra/config"
	"github.com/felipeguilhermefs/selene/infra/errors"
	"github.com/felipeguilhermefs/selene/repositories"
	"github.com/felipeguilhermefs/selene/services"
	"github.com/felipeguilhermefs/selene/view"
)

// Server represents all insfrastructure used in this server app
type Server struct {
	repositories *repositories.Repositories
	services     *services.Services
	controllers  *controllers.Controllers
	server       *http.Server
	sessionStore *sessions.CookieStore
}

// Start start listening and serving requests
func (s *Server) Start() error {
	log.Printf("Server started at %v...\n", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *Server) handleBooksPage() http.HandlerFunc {
	type book struct {
		ID     uint
		Title  string
		Author string
		Tags   string
	}

	data := view.Data{
		User: "someone",
		Yield: []book{
			{1, "The Hobbit", "JRR Tolkien", "adventure, fantasy"},
			{2, "Do Androids Dream of Electric Sheep?", "Philip K. Dick", "sci-fi, philosophical"},
			{3, "1984", "George Orwell", "dystopian, political fiction"},
		},
	}

	page := view.NewView("books")
	return func(w http.ResponseWriter, r *http.Request) {
		page.Render(w, r, &data)
	}
}

func (s *Server) handleBookPage() http.HandlerFunc {
	type book struct {
		ID       uint
		Title    string
		Author   string
		Comments string
		Tags     string
	}

	data := view.Data{
		User: "someone",
		Yield: book{
			ID:       2,
			Title:    "Do Androids Dream of Electric Sheep?",
			Author:   "Philip K. Dick",
			Comments: "quero muito ler esse",
			Tags:     "sci-fi, philosophical",
		},
	}

	page := view.NewView("book")
	return func(w http.ResponseWriter, r *http.Request) {
		page.Render(w, r, &data)
	}
}

func (s *Server) handleNewBookPage() http.HandlerFunc {
	type book struct {
		Title    string
		Author   string
		Comments string
		Tags     string
	}

	data := view.Data{
		User: "someone",
		Yield: book{
			Title:    "American Gods",
			Author:   "Neil Gaiman",
			Comments: "opa",
			Tags:     "fantasy",
		},
	}

	page := view.NewView("new_book")
	return func(w http.ResponseWriter, r *http.Request) {
		page.Render(w, r, &data)
	}
}

// NewServer creates a new server instance
func NewServer(cfg *config.Config) (*Server, error) {
	router := mux.NewRouter()

	repos, err := repositories.NewRepositories(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "Creating repositories")
	}

	if err := repos.AutoMigrate(); err != nil {
		return nil, errors.Wrap(err, "Migrating repositories")
	}

	srvcs := services.NewServices(cfg, repos)

	ctrls := controllers.NewControllers(router, srvcs)

	sessionStore := sessions.NewCookieStore(
		[]byte(cfg.Sec.Session.AuthKey),
		[]byte(cfg.Sec.Session.CryptoKey),
	)

	sessionStore.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
	}
	sessionStore.MaxAge(cfg.Sec.Session.TTL)

	s := Server{
		repositories: repos,
		services:     srvcs,
		controllers:  ctrls,
		server: &http.Server{
			Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
			ReadTimeout:  cfg.Server.ReadTimeout(),
			WriteTimeout: cfg.Server.WriteTimeout(),
			IdleTimeout:  cfg.Server.IdleTimeout(),
			Handler:      router,
		},
		sessionStore: sessionStore,
	}

	router.HandleFunc("/books", s.handleBooksPage()).Methods("GET")
	router.HandleFunc("/books/{id:[0-9]+}", s.handleBookPage()).Methods("GET")
	router.HandleFunc("/books/new", s.handleNewBookPage()).Methods("GET")

	return &s, nil
}
