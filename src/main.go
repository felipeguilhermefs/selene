package main

import (
	"log"

	"github.com/felipeguilhermefs/selene/handlers"
	"github.com/felipeguilhermefs/selene/infra/config"
	"github.com/felipeguilhermefs/selene/infra/database"
	"github.com/felipeguilhermefs/selene/infra/session"
	"github.com/felipeguilhermefs/selene/infrastructure/middleware"
	"github.com/felipeguilhermefs/selene/infrastructure/router"
	"github.com/felipeguilhermefs/selene/infrastructure/server"
	"github.com/felipeguilhermefs/selene/middlewares"
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

	sessionStore := session.NewCookieStore(&cfg.Sec.Session)

	repos := repositories.New(db, sessionStore)
	if err := repos.AutoMigrate(); err != nil {
		return err
	}

	srvcs := services.New(&cfg.Sec.Password, repos)

	views := view.NewViews()

	hdlrs := handlers.New(srvcs, views)

	mdw := middlewares.New(cfg.Sec.CSRF, srvcs.Auth, hdlrs.NotAuthentic)

	routes := []router.Route{
		{Method: "GET", Path: "/signup", Handler: hdlrs.SignupPage},
		{Method: "POST", Path: "/signup", Handler: hdlrs.Signup},
		{Method: "GET", Path: "/login", Handler: hdlrs.LoginPage},
		{Method: "POST", Path: "/login", Handler: hdlrs.Login},
		{Method: "POST", Path: "/logout", Handler: mdw.Login(hdlrs.Logout)},
		{Method: "GET", Path: "/books", Handler: mdw.Login(hdlrs.BooksPage)},
		{Method: "GET", Path: "/books/new", Handler: mdw.Login(hdlrs.NewBookPage)},
		{Method: "POST", Path: "/books/new", Handler: mdw.Login(hdlrs.NewBook)},
		{Method: "GET", Path: "/books/{id:[0-9]+}", Handler: mdw.Login(hdlrs.BookPage)},
		{Method: "POST", Path: "/books/{id:[0-9]+}/edit", Handler: mdw.Login(hdlrs.EditBook)},
		{Method: "POST", Path: "/books/{id:[0-9]+}/delete", Handler: mdw.Login(hdlrs.DeleteBook)},
	}

	mdws := []middleware.Middleware{mdw.CSRF, mdw.SecHeaders}

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
