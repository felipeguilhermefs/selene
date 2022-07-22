package view

import (
	"bytes"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/csrf"
)

const (
	baseLayout               = "base"
	ErrNoCSRFField CSRFError = "No CSRF field implemented"
)

// View represents a page that renders (lazily) from a template
type View struct {
	name      string
	once      sync.Once
	layouts   string
	templates fs.FS
	tpl       *template.Template
	tplerr    error
}

// Render will render and enrich a view template with provided data
func (v *View) Render(w http.ResponseWriter, r *http.Request, data *Data) {
	v.once.Do(func() { v.parse() })

	if v.tplerr != nil {
		log.Println(v.tplerr)
		http.Error(w, defaultErrorMessage, http.StatusInternalServerError)
		return
	}

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
	v.tpl, v.tplerr = template.New("").
		Funcs(v.csrfTag(nil)).
		ParseFS(v.templates, v.name, v.layouts)
}

func (v *View) csrfTag(r *http.Request) template.FuncMap {
	custom := template.FuncMap{}

	custom[csrf.TemplateTag] = func() (template.HTML, error) {
		if r == nil {
			return "", ErrNoCSRFField
		}

		return csrf.TemplateField(r), nil
	}

	return custom
}

type CSRFError string

func (e CSRFError) Error() string {
	return string(e)
}
