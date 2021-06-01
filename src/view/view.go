package view

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gorilla/csrf"

	"github.com/felipeguilhermefs/selene/infra/errors"
)

const baseLayout = "base"

// View represents a page that renders (lazily) from a template
type View struct {
	name   string
	once   sync.Once
	tpl    *template.Template
	tplerr error
	csp    string
	static StaticData
}

// Render will render and enrich a view template with provided data
func (v *View) Render(w http.ResponseWriter, r *http.Request, data *Data) {
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	v.once.Do(func() {
		v.parse()

		v.csp = buildCSP(allowedScripts, allowedStyles)

		v.static = StaticData{
			Scripts: allowedScripts,
			Styles:  allowedStyles,
		}
	})

	if v.tplerr != nil {
		log.Println(v.tplerr)
		http.Error(w, defaultErrorMessage, http.StatusInternalServerError)
		return
	}

	v.setContentSecurityPolicy(w, data)

	tpl := v.tpl.Funcs(v.csrfTag(r))

	var buf bytes.Buffer
	if err := tpl.ExecuteTemplate(&buf, baseLayout, data); err != nil {
		log.Println(err)
		http.Error(w, defaultErrorMessage, http.StatusInternalServerError)
		return
	}

	io.Copy(w, &buf)
}

func (v *View) parse() {
	templateFile := fmt.Sprintf("view/templates/%s.gohtml", v.name)

	layoutFiles, err := filepath.Glob("view/templates/layouts/*.gohtml")
	if err != nil {
		v.tplerr = err
	}

	templateFiles := append([]string{templateFile}, layoutFiles...)

	v.tpl, v.tplerr = template.New("").
		Funcs(v.csrfTag(nil)).
		ParseFiles(templateFiles...)
}

func (v *View) csrfTag(r *http.Request) template.FuncMap {
	custom := template.FuncMap{}

	custom[csrf.TemplateTag] = func() (template.HTML, error) {
		if r == nil {
			return "", errors.ErrNoCSRFField
		}

		return csrf.TemplateField(r), nil
	}

	return custom
}

func (v *View) setContentSecurityPolicy(w http.ResponseWriter, data *Data) {
	data.Static = v.static

	w.Header().Set("Content-Security-Policy", v.csp)
}

func buildCSP(scripts []Dependency, styles []Dependency) string {
	csp := []string{
		"default-src 'none'",
		"base-uri 'self'",
		"form-action 'self'",
		"frame-ancestors 'none'",
		"upgrade-insecure-requests",
	}

	cspScripts := "script-src "
	for _, script := range scripts {
		cspScripts += script.URL + " "
	}

	cspStyles := "style-src "
	for _, style := range styles {
		cspStyles += style.URL + " "
	}

	csp = append(csp, cspScripts, cspStyles)

	return strings.Join(csp, ";")
}
