package main

import (
	"log"

	"github.com/felipeguilhermefs/selene/handlers"
	"github.com/felipeguilhermefs/selene/infra/config"
	"github.com/felipeguilhermefs/selene/infra/database"
	"github.com/felipeguilhermefs/selene/infra/session"
	"github.com/felipeguilhermefs/selene/infrastructure/server"
	"github.com/felipeguilhermefs/selene/middlewares"
	"github.com/felipeguilhermefs/selene/repositories"
	"github.com/felipeguilhermefs/selene/router"
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

	r := router.New(mdw, hdlrs)

	s := server.New(&cfg.Server, r)

	log.Printf("Server started at %v...\n", s.Addr)
	return s.ListenAndServe()
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
