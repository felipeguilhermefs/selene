package hsts

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	hsts = "Strict-Transport-Security"
)

type Config interface {
	IncludeSubDomains() bool
	MaxAge() int
	Preload() bool
}

func New(cfg Config) func(next http.Handler) http.Handler {
	hstsValue := build(cfg)

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set(hsts, hstsValue)

			next.ServeHTTP(w, r)
		})
	}
}

func build(cfg Config) string {
	var rules []string

	if cfg.MaxAge() > 0 {
		rules = append(rules, fmt.Sprintf("max-age=%d", cfg.MaxAge()))
	}

	if cfg.IncludeSubDomains() {
		rules = append(rules, "includeSubDomains")
	}

	if cfg.Preload() {
		rules = append(rules, "preload")
	}

	return strings.Join(rules, ";")
}
