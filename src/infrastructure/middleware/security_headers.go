package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	referrer = "Referrer-Policy"
	csp      = "Content-Security-Policy"
	hsts     = "Strict-Transport-Security"
)

type SecHeaderConfig struct {
	CSP            CSPConfig
	HSTS           HSTSConfig
	ReferrerPolicy string
}

type CSPConfig struct {
	BaseURI         string
	DefaultSrc      string
	FormAction      string
	FrameAncestors  string
	StyleSrc        []string
	ScriptSrc       []string
	UpgradeInsecure bool
}

type HSTSConfig struct {
	IncludeSubDomains bool
	MaxAge            int
	Preload           bool
}

func NewSecHeaders(cfg *SecHeaderConfig) Middleware {
	cspValue := cfg.CSP.build()
	hstsValue := cfg.HSTS.build()

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set(csp, cspValue)
			w.Header().Set(referrer, cfg.ReferrerPolicy)
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

func (c *CSPConfig) build() string {
	var rules []string

	if defaultSrc := strings.TrimSpace(c.DefaultSrc); defaultSrc != "" {
		rules = append(rules, fmt.Sprintf("default-src %s", defaultSrc))
	}

	if baseURI := strings.TrimSpace(c.BaseURI); baseURI != "" {
		rules = append(rules, fmt.Sprintf("base-uri %s", baseURI))
	}

	if formAction := strings.TrimSpace(c.FormAction); formAction != "" {
		rules = append(rules, fmt.Sprintf("form-action %s", formAction))
	}

	if frameAncestors := strings.TrimSpace(c.FrameAncestors); frameAncestors != "" {
		rules = append(rules, fmt.Sprintf("frame-ancestors %s", frameAncestors))
	}

	if c.UpgradeInsecure {
		rules = append(rules, "upgrade-insecure-requests")
	}

	if len(c.ScriptSrc) > 0 {
		scriptSrc := fmt.Sprintf("script-src %s", strings.Join(c.ScriptSrc, " "))
		rules = append(rules, scriptSrc)
	}

	if len(c.StyleSrc) > 0 {
		styleSrc := fmt.Sprintf("style-src %s", strings.Join(c.StyleSrc, " "))
		rules = append(rules, styleSrc)
	}

	return strings.Join(rules, ";")
}
