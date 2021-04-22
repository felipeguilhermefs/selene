package main

import (
	"net/http"

	"github.com/gorilla/mux"

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
	router       *mux.Router
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

// NewServer creates a new server instance
func NewServer(cfg *config.Config) (*Server, error) {
	router := mux.NewRouter()

	repos, err := repositories.NewRepositories(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "Creating repositories")
	}

	srvcs := services.NewServices(cfg, repos)

	ctrls := controllers.NewControllers(srvcs)
	ctrls.RegisterRoutes(router)

	s := Server{
		repositories: repos,
		services:     srvcs,
		controllers:  ctrls,
		router:       router,
	}

	router.HandleFunc("/books", s.handleBooksPage()).Methods("GET")
	router.HandleFunc("/books/{id:[0-9]+}", s.handleBookPage()).Methods("GET")
	router.HandleFunc("/books/new", s.handleNewBookPage()).Methods("GET")
	router.HandleFunc("/login", s.handleLoginPage()).Methods("GET")

	return &s, nil
}
