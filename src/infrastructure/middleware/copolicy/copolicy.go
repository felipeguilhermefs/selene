package copolicy

import "net/http"

const (
	coop = "Cross-Origin-Opener-Policy"
	coep = "Cross-Origin-Embedder-Policy"
	corp = "Cross-Origin-Resource-Policy"
)

type Config struct {
	Embedder string
	Opener   string
	Resource string
}

func New(cfg Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set(coep, cfg.Embedder)
			w.Header().Set(coop, cfg.Opener)
			w.Header().Set(corp, cfg.Resource)

			next.ServeHTTP(w, r)
		})
	}
}
