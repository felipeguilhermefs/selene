package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	hsts = "Strict-Transport-Security"
)

type SecHeaderConfig struct {
	HSTS           HSTSConfig
}

type HSTSConfig struct {
	IncludeSubDomains bool
	MaxAge            int
	Preload           bool
}

func NewSecHeaders(cfg *SecHeaderConfig) Middleware {
	hstsValue := cfg.HSTS.build()

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set(hsts, hstsValue)

			next.ServeHTTP(w, r)
		})
	}
}

func (c *HSTSConfig) build() string {
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
