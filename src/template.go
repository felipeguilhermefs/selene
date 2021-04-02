package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

const (
	baseLayout          = "base"
	defaultErrorMessage = "Something whent wrong. Please contact us if this error persists"
)

// TemplateData dynamic data to enrich templates
type TemplateData struct {
	User  interface{}
	Yield interface{}
}

// TemplateDataFetcher is a function used to fetch dynamic data to be
// injected into a template before its render
type TemplateDataFetcher func(r *http.Request) (*TemplateData, error)

// HandleTemplate lazily initialize a template with the given name and
// renders it with dynamic data retrieved from fetcher.
//
// name argument should be a file name (without extension) in templates directory.
//    EX: HandleTemplate("books", someDataFetcher) will use "templates/books.gohtml"
func HandleTemplate(templateName string, fetcher TemplateDataFetcher) http.HandlerFunc {
	var (
		once   sync.Once
		tpl    *template.Template
		tplerr error
	)

	return func(w http.ResponseWriter, r *http.Request) {
		once.Do(func() {
			tpl, tplerr = parseTemplate(templateName)
		})

		if tplerr != nil {
			log.Println(WrapError(tplerr, "Template parsing"))
			http.Error(w, defaultErrorMessage, http.StatusInternalServerError)
			return
		}

		data, err := fetcher(r)
		if err != nil {
			log.Println(WrapError(err, "Template data fetching"))
			http.Error(w, defaultErrorMessage, http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "text/html")

		var buf bytes.Buffer
		if err := tpl.ExecuteTemplate(&buf, baseLayout, data); err != nil {
			log.Println(WrapError(err, "Template rendering"))
			http.Error(w, defaultErrorMessage, http.StatusInternalServerError)
			return
		}

		io.Copy(w, &buf)
	}
}

func parseTemplate(templateName string) (*template.Template, error) {
	templateFile := fmt.Sprintf("templates/%s.gohtml", templateName)

	layoutFiles, err := filepath.Glob("templates/layouts/*.gohtml")
	if err != nil {
		return nil, WrapError(err, "Finding layout files")
	}

	templateFiles := append([]string{templateFile}, layoutFiles...)

	tpl, err := template.ParseFiles(templateFiles...)
	if err != nil {
		return nil, WrapError(err, "Parsing files found")
	}

	return tpl, nil
}
