package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/felipeguilhermefs/selene/infra/config"
	"github.com/felipeguilhermefs/selene/infra/database"
	"github.com/felipeguilhermefs/selene/infra/errors"
	"github.com/felipeguilhermefs/selene/view"
)

// Server represents all insfrastructure used in this server app
type Server struct {
	db     *gorm.DB
	router *mux.Router
}

// ServeHTTP just delegates to router so we can start Server in http.ListenAndServe
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
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

func (s *Server) handleLoginPage() http.HandlerFunc {
	type login struct {
		Email string
	}

	data := view.Data{
		Yield: login{
			Email: "king@nnt.leo",
		},
	}

	page := view.NewView("login")
	return func(w http.ResponseWriter, r *http.Request) {
		page.Render(w, r, &data)
	}
}

func (s *Server) handleSignupPage() http.HandlerFunc {
	type signup struct {
		Name  string
		Email string
	}

	data := view.Data{
		Yield: signup{
			Name:  "Harlequin",
			Email: "king@nnt.leo",
		},
	}

	page := view.NewView("signup")
	return func(w http.ResponseWriter, r *http.Request) {
		page.Render(w, r, &data)
	}
}

// NewServer creates a new server instance
func NewServer(cfg *config.Config) (*Server, error) {
	db, err := database.ConnectPostgres(&cfg.Postgres)
	if err != nil {
		return nil, errors.Wrap(err, "Connecting to Postgres")
	}

	router := mux.NewRouter()

	s := Server{
		db:     db,
		router: router,
	}

	router.HandleFunc("/books", s.handleBooksPage()).Methods("GET")
	router.HandleFunc("/books/{id:[0-9]+}", s.handleBookPage()).Methods("GET")
	router.HandleFunc("/books/new", s.handleNewBookPage()).Methods("GET")
	router.HandleFunc("/login", s.handleLoginPage()).Methods("GET")
	router.HandleFunc("/signup", s.handleSignupPage()).Methods("GET")

	return &s, nil
}
