package csp

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	csp = "Content-Security-Policy"
)

type Config interface {
	BaseURI() string
	DefaultSrc() string
	FormAction() string
	FrameAncestors() string
	StyleSrc() []string
	ScriptSrc() []string
	UpgradeInsecure() bool
}

func New(cfg Config) func(next http.Handler) http.Handler {
	cspValue := build(cfg)

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set(csp, cspValue)

			next.ServeHTTP(w, r)
		})
	}
}

func build(cfg Config) string {
	var rules []string

	if defaultSrc := strings.TrimSpace(cfg.DefaultSrc()); defaultSrc != "" {
		rules = append(rules, fmt.Sprintf("default-src %s", defaultSrc))
	}

	if baseURI := strings.TrimSpace(cfg.BaseURI()); baseURI != "" {
		rules = append(rules, fmt.Sprintf("base-uri %s", baseURI))
	}

	if formAction := strings.TrimSpace(cfg.FormAction()); formAction != "" {
		rules = append(rules, fmt.Sprintf("form-action %s", formAction))
	}

	if frameAncestors := strings.TrimSpace(cfg.FrameAncestors()); frameAncestors != "" {
		rules = append(rules, fmt.Sprintf("frame-ancestors %s", frameAncestors))
	}

	if cfg.UpgradeInsecure() {
		rules = append(rules, "upgrade-insecure-requests")
	}

	if len(cfg.ScriptSrc()) > 0 {
		scriptSrc := fmt.Sprintf("script-src %s", strings.Join(cfg.ScriptSrc(), " "))
		rules = append(rules, scriptSrc)
	}

	if len(cfg.StyleSrc()) > 0 {
		styleSrc := fmt.Sprintf("style-src %s", strings.Join(cfg.StyleSrc(), " "))
		rules = append(rules, styleSrc)
	}

	return strings.Join(rules, ";")
}
