package hsts

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	hsts = "Strict-Transport-Security"
)

type Config struct {
	IncludeSubDomains bool
	MaxAge            int
	Preload           bool
}

func New(cfg *Config) func(next http.Handler) http.Handler {
	hstsValue := cfg.build()

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set(hsts, hstsValue)

			next.ServeHTTP(w, r)
		})
	}
}

func (c *Config) build() string {
	var rules []string

	if c.MaxAge > 0 {
		rules = append(rules, fmt.Sprintf("max-age=%d", c.MaxAge))
	}

	if c.IncludeSubDomains {
		rules = append(rules, "includeSubDomains")
	}

	if c.Preload {
		rules = append(rules, "preload")
	}

	return strings.Join(rules, ";")
}
