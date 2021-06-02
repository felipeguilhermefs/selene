package middlewares

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	referrerHeader = "Referrer-Policy"
	cspHeader      = "Content-Security-Policy"

	jquery       = "https://code.jquery.com/jquery-3.5.1.slim.min.js"
	bootstrapJS  = "https://cdn.jsdelivr.net/npm/bootstrap@4.6.0/dist/js/bootstrap.bundle.min.js"
	bootstrapCSS = "https://cdn.jsdelivr.net/npm/bootstrap@4.6.0/dist/css/bootstrap.min.css"
)

func newSecHeaderMiddleware() Middleware {
	cspValue := buildCSP()

	return func(next http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set(cspHeader, cspValue)
			w.Header().Set(referrerHeader, "no-referrer")

			next(w, r)
		}
	}
}

func buildCSP() string {
	scriptSrc := fmt.Sprintf("script-src %s %s", jquery, bootstrapJS)
	styleSrc := fmt.Sprintf("style-src %s", bootstrapCSS)

	csp := []string{
		"default-src 'none'",
		"base-uri 'self'",
		"form-action 'self'",
		"frame-ancestors 'none'",
		"upgrade-insecure-requests",
		scriptSrc,
		styleSrc,
	}

	return strings.Join(csp, ";")
}
