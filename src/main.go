package main

import (
	"embed"
	"log"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/felipeguilhermefs/selene/boundary/postgres"
	"github.com/felipeguilhermefs/selene/core/bookshelf"
	"github.com/felipeguilhermefs/selene/handlers"
	"github.com/felipeguilhermefs/selene/infrastructure/config"
	"github.com/felipeguilhermefs/selene/infrastructure/database"
	"github.com/felipeguilhermefs/selene/infrastructure/middleware/auth"
	"github.com/felipeguilhermefs/selene/infrastructure/middleware/csrf"
	"github.com/felipeguilhermefs/selene/infrastructure/middleware/hsts"
	htmlMiddleware "github.com/felipeguilhermefs/selene/infrastructure/middleware/html"
	"github.com/felipeguilhermefs/selene/infrastructure/middleware/policy"
	"github.com/felipeguilhermefs/selene/infrastructure/router"
	"github.com/felipeguilhermefs/selene/infrastructure/server"
	"github.com/felipeguilhermefs/selene/infrastructure/session"
	"github.com/felipeguilhermefs/selene/services"
	"github.com/felipeguilhermefs/selene/view"
)

//go:embed templates
var templates embed.FS

func run() error {
	cfg := config.New()

	db, err := database.Connect(cfg)
	if err != nil {
		return err
	}

	sessionStore := session.NewStore(cfg)

	if err := postgres.RunMigrations(db); err != nil {
		return err
	}

	userRepository := &postgres.PostgresUserRepository{
		DB:         db,
		EmailRegex: regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,16}$`),
	}

	srvcs := services.New(cfg, userRepository, sessionStore)

	views := view.NewViews(templates)

	authenticated := auth.New(srvcs.Auth)
	html := htmlMiddleware.New()

	bookControl := &bookshelf.BookControl{
		BookRepository: &postgres.PostgresBookRepository{
			DB: db,
		},
	}

	hdlrs := handlers.New(srvcs, views, bookControl)

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
		csrf.New(cfg),
		policy.Policy,
		hsts.HSTS,
	}

	r := router.New(routes, mdws, hdlrs.NotFound)

	exiting := make(chan error, 1)

	s := server.New(cfg, r)
	go func() {
		exiting <- s.Serve()
	}()

	go func() {
		gracefulShutdown := make(chan os.Signal, 1)
		signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGTERM)

		log.Printf("Received %s, gracefully shutting down...", <-gracefulShutdown)

		exiting <- s.Shutdown()
	}()

	return <-exiting

}

func main() {
	if err := run(); err != nil {
		log.Fatalln("Fatal Error", err)
	}
}
