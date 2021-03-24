package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"
)

const (
	baseLayout = "base"
	defaultErrorMessage = "Something whent wrong. Please contact us if this error persists"
)

// NewView creates a new instance of a view from its template file
func NewView(templateName string) (*View, error) {
	templateFiles, err := findTemplateFiles(templateName)
	if err != nil {
		return nil, err
	}

	t, err := template.New("").ParseFiles(templateFiles...)
	if err != nil {
		return nil, err
	}

	return &View{
		template: t,
	}, nil
}

// View http.Handler that holds reference to template files and control its rendering
type View struct {
	template *template.Template
}

// ServeHTTP render the view with predefined layout
func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.Render(w, r, nil)
}

// Render write template with view data to response writer
func (v *View) Render(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "text/html")

	var buf bytes.Buffer
	if err := v.template.ExecuteTemplate(&buf, baseLayout, nil); err != nil {
		log.Println(err)
		http.Error(w, defaultErrorMessage, http.StatusInternalServerError)
		return
	}

	io.Copy(w, &buf)
}

func findTemplateFiles(templateName string) ([]string, error) {
	templateFile := fmt.Sprintf("templates/%s.gohtml", templateName)

	layoutFiles, err := filepath.Glob("templates/layouts/*.gohtml")
	if err != nil {
		return nil, err
	}

	templateFiles := append([]string{templateFile}, layoutFiles...)

	return templateFiles, nil
}
