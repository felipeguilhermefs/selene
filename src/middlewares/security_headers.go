package middlewares

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	referrer = "Referrer-Policy"
	csp      = "Content-Security-Policy"
	coop     = "Cross-Origin-Opener-Policy"
	coep     = "Cross-Origin-Embedder-Policy"
	corp     = "Cross-Origin-Resource-Policy"

	jquery       = "https://code.jquery.com/jquery-3.5.1.slim.min.js"
	bootstrapJS  = "https://cdn.jsdelivr.net/npm/bootstrap@4.6.0/dist/js/bootstrap.bundle.min.js"
	bootstrapCSS = "https://cdn.jsdelivr.net/npm/bootstrap@4.6.0/dist/css/bootstrap.min.css"
)

func newSecHeaderMiddleware() Middleware {
	cspValue := buildCSP()

	return func(next http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set(csp, cspValue)
			w.Header().Set(referrer, "no-referrer")
			w.Header().Set(coop, "same-origin")
			w.Header().Set(coep, "require-corp")
			w.Header().Set(corp, "same-origin")

			next(w, r)
		}
	}
}

func buildCSP() string {
	scriptSrc := fmt.Sprintf("script-src %s %s", jquery, bootstrapJS)
	styleSrc := fmt.Sprintf("style-src %s", bootstrapCSS)

	rules := []string{
		"default-src 'none'",
		"base-uri 'self'",
		"form-action 'self'",
		"frame-ancestors 'none'",
		"upgrade-insecure-requests",
		scriptSrc,
		styleSrc,
	}

	return strings.Join(rules, ";")
}
