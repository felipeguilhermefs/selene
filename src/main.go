package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/felipeguilhermefs/selene/boundary/postgres"
	"github.com/felipeguilhermefs/selene/core/auth"
	"github.com/felipeguilhermefs/selene/core/bookshelf"
	"github.com/felipeguilhermefs/selene/handlers"
	"github.com/felipeguilhermefs/selene/infrastructure/config"
	authMiddleware "github.com/felipeguilhermefs/selene/infrastructure/middleware/auth"
	"github.com/felipeguilhermefs/selene/infrastructure/middleware/csrf"
	"github.com/felipeguilhermefs/selene/infrastructure/middleware/hsts"
	htmlMiddleware "github.com/felipeguilhermefs/selene/infrastructure/middleware/html"
	"github.com/felipeguilhermefs/selene/infrastructure/middleware/policy"
	"github.com/felipeguilhermefs/selene/infrastructure/router"
	"github.com/felipeguilhermefs/selene/infrastructure/server"
	"github.com/felipeguilhermefs/selene/infrastructure/session"
	"github.com/felipeguilhermefs/selene/templates"
	"github.com/felipeguilhermefs/selene/view"
)

func run() error {
	cfg := config.New()

	pg, err := postgres.New(cfg)
	if err != nil {
		return err
	}

	sessionStore := session.NewStore(cfg)

	if err := pg.RunMigrations(); err != nil {
		return err
	}

	authControl := &auth.AuthControl{
		UserRepository:  pg.UserRepository,
		EmailNormalizer: auth.EmailNormalizer{},
		PasswordControl: auth.NewPasswordControl(cfg),
	}

	views := view.NewViews(templates.FS)

	authenticated := authMiddleware.New(authControl, sessionStore)
	html := htmlMiddleware.New()

	bookshelfControl := &bookshelf.BookshelfControl{
		BookRepository: pg.BookRepository,
	}

	hdlrs := handlers.New(views, bookshelfControl, authControl, authControl, sessionStore)

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
