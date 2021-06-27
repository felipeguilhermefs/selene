package policy

import "net/http"

const (
	coop     = "Cross-Origin-Opener-Policy"
	coep     = "Cross-Origin-Embedder-Policy"
	corp     = "Cross-Origin-Resource-Policy"
	referrer = "Referrer-Policy"
)

type Config interface {
	Embedder() string
	Opener() string
	Referrer() string
	Resource() string
}

func New(cfg Config) func(next http.Handler) http.Handler {
	embedder := cfg.Embedder()
	opener := cfg.Opener()
	referrer := cfg.Referrer()
	resource := cfg.Resource()
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set(coep, embedder)
			w.Header().Set(coop, opener)
			w.Header().Set(corp, resource)
			w.Header().Set(referrer, referrer)

			next.ServeHTTP(w, r)
		})
	}
}
