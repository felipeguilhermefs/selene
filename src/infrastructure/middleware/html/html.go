package html

import (
	"net/http"
	"strings"
)

func New() func(next http.Handler) http.Handler {
	cspValue := buildCSP()

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Header().Set("Content-Security-Policy", cspValue)

			next.ServeHTTP(w, r)
		})
	}
}

// CSP helps reduce XSS attack surface, and other types of data injection.
// More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP
func buildCSP() string {
	rules := []string{
		"default-src 'none'",
		"base-uri 'self'",
		"form-action 'self'",
		"frame-ancestors 'none'",
		"upgrade-insecure-requests",
		"script-src https://code.jquery.com/jquery-3.5.1.slim.min.js https://cdn.jsdelivr.net/npm/bootstrap@4.6.0/dist/js/bootstrap.bundle.min.js",
		"style-src https://cdn.jsdelivr.net/npm/bootstrap@4.6.0/dist/css/bootstrap.min.css",
	}

	return strings.Join(rules, ";")
}
