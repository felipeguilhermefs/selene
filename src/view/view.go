package view

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"
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
}

// Render will render and enrich a view template with provided data
func (v *View) Render(w http.ResponseWriter, r *http.Request, data *Data) {
	w.Header().Set("Content-Type", "text/html")

	v.once.Do(func() {
		v.csp = buildCSP(allowedScripts)
		v.parse()
	})

	if v.tplerr != nil {
		log.Println(v.tplerr)
		http.Error(w, defaultErrorMessage, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Security-Policy", v.csp)
	data.Static.Scripts = allowedScripts

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

func buildCSP(scripts []Script) string {
	csp := "script-src "

	for _, script := range scripts {
		csp += script.URL + " "
	}

	return csp
}
