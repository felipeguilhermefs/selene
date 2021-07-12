package main

import (
	"log"

	"github.com/felipeguilhermefs/selene/handlers"
	"github.com/felipeguilhermefs/selene/infra/config"
	"github.com/felipeguilhermefs/selene/infra/database"
	"github.com/felipeguilhermefs/selene/infrastructure/middleware/auth"
	"github.com/felipeguilhermefs/selene/infrastructure/middleware/csrf"
	"github.com/felipeguilhermefs/selene/infrastructure/middleware/hsts"
	htmlMiddleware "github.com/felipeguilhermefs/selene/infrastructure/middleware/html"
	"github.com/felipeguilhermefs/selene/infrastructure/middleware/policy"
	"github.com/felipeguilhermefs/selene/infrastructure/router"
	"github.com/felipeguilhermefs/selene/infrastructure/server"
	"github.com/felipeguilhermefs/selene/infrastructure/session"
	"github.com/felipeguilhermefs/selene/repositories"
	"github.com/felipeguilhermefs/selene/services"
	"github.com/felipeguilhermefs/selene/view"
)

func run() error {
	cfg, err := config.LoadFromFile("config.json")
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	db, err := database.ConnectPostgres(&cfg.DB)
	if err != nil {
		return err
	}

	sessionStore := session.NewStore(&cfg.Sec.Session)

	repos := repositories.New(db)
	if err := repos.AutoMigrate(); err != nil {
		return err
	}

	srvcs := services.New(&cfg.Sec.Password, repos, sessionStore)

	views := view.NewViews()

	authenticated := auth.New(srvcs.Auth)
	html := htmlMiddleware.New()

	hdlrs := handlers.New(srvcs, views)

	routes := []router.Route{
		{Method: "GET", Path: "/signup", Handler: html(hdlrs.SignupPage)},
		{Method: "POST", Path: "/signup", Handler: hdlrs.Signup},
		{Method: "GET", Path: "/login", Handler: html(hdlrs.LoginPage)},
		{Method: "POST", Path: "/login", Handler: hdlrs.Login},
		{Method: "POST", Path: "/logout", Handler: authenticated(hdlrs.Logout)},
		{Method: "GET", Path: "/books", Handler: authenticated(html(hdlrs.BooksPage))},
		{Method: "GET", Path: "/books/new", Handler: authenticated(html(hdlrs.NewBookPage))},
		{Method: "POST", Path: "/books/new", Handler: authenticated(hdlrs.NewBook)},
		{Method: "GET", Path: "/books/{id:[0-9]+}", Handler: authenticated(html(hdlrs.BookPage))},
		{Method: "POST", Path: "/books/{id:[0-9]+}/edit", Handler: authenticated(hdlrs.EditBook)},
		{Method: "POST", Path: "/books/{id:[0-9]+}/delete", Handler: authenticated(hdlrs.DeleteBook)},
	}

	mdws := []router.Middleware{
		csrf.New(&cfg.Sec),
		policy.New(&cfg.Sec),
		hsts.New(&cfg.Sec),
	}

	r := router.New(routes, mdws, hdlrs.NotFound)

	s := server.New(&cfg.Server, r)

	log.Printf("Server started at %v...\n", s.Addr)
	return s.ListenAndServe()
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
