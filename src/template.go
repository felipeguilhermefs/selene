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

// TemplateDataFetcher is a function used to fetch dynamic data to be
// injected into a template before its render
type TemplateDataFetcher func (r *http.Request) (interface{}, error)

// HandleTemplate lazily initialize a template with the given name and
// renders it with dynamic data retrieved from fetcher.
//
// name argument should be a file name (without extension) in templates directory.
//    EX: HandleTemplate("books", someDataFetcher) will use "templates/books.gohtml"
func HandleTemplate(name string, fetcher TemplateDataFetcher) http.HandlerFunc {
	var (
		once   sync.Once
		tpl    *template.Template
		tplerr error
	)

	return func(w http.ResponseWriter, r *http.Request) {
		once.Do(func(){
			templateFiles, err := findTemplateFiles(name)
			if err != nil {
				tpl, tplerr = nil, err
				return
			}

			tpl, tplerr = template.ParseFiles(templateFiles...)
		})

		if tplerr != nil {
			log.Println(tplerr)
			http.Error(w, defaultErrorMessage, http.StatusInternalServerError)
			return
		}

		data, err := fetcher(r)
		if err != nil {
			log.Println(err)
			http.Error(w, defaultErrorMessage, http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "text/html")

		var buf bytes.Buffer
		if err := tpl.ExecuteTemplate(&buf, baseLayout, data); err != nil {
			log.Println(err)
			http.Error(w, defaultErrorMessage, http.StatusInternalServerError)
			return
		}

		io.Copy(w, &buf)
	}
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
