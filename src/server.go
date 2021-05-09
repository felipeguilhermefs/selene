package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/felipeguilhermefs/selene/handlers"
	"github.com/felipeguilhermefs/selene/infra/config"
	"github.com/felipeguilhermefs/selene/infra/database"
	"github.com/felipeguilhermefs/selene/infra/errors"
	"github.com/felipeguilhermefs/selene/infra/session"
	"github.com/felipeguilhermefs/selene/repositories"
	"github.com/felipeguilhermefs/selene/services"
	"github.com/felipeguilhermefs/selene/view"
)

// Server represents all insfrastructure used in this server app
type Server struct {
	repositories *repositories.Repositories
	services     *services.Services
	server       *http.Server
}

// Start start listening and serving requests
func (s *Server) Start() error {
	log.Printf("Server started at %v...\n", s.server.Addr)
	return s.server.ListenAndServe()
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

	s := Server{
		repositories: repos,
		services:     srvcs,
		server: &http.Server{
			Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
			ReadTimeout:  cfg.Server.ReadTimeout(),
			WriteTimeout: cfg.Server.WriteTimeout(),
			IdleTimeout:  cfg.Server.IdleTimeout(),
			Handler:      router,
		},
	}

	loginView := view.NewView("login")
	signupView := view.NewView("signup")

	router.HandleFunc("/signup", handlers.HandleSignupPage(signupView)).Methods("GET")
	router.HandleFunc("/signup", handlers.HandleSignup(signupView, srvcs.Auth)).Methods("POST")
	router.HandleFunc("/login", handlers.HandleLoginPage(loginView)).Methods("GET")
	router.HandleFunc("/login", handlers.HandleLogin(loginView, srvcs.Auth)).Methods("POST")
	router.HandleFunc("/logout", handlers.HandleLogout(srvcs.Auth)).Methods("POST")

	booksView := view.NewView("books")

	router.HandleFunc("/books", handlers.HandleBooksPage(booksView, srvcs.Auth, srvcs.Book)).Methods("GET")
	router.HandleFunc("/books/{id:[0-9]+}", s.handleBookPage()).Methods("GET")
	router.HandleFunc("/books/new", s.handleNewBookPage()).Methods("GET")

	return &s, nil
}
