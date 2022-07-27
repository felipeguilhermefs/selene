package handlers

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/schema"
)

func parseForm(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	return parseValues(r.PostForm, dst)
}

func parseURLParams(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	return parseValues(r.Form, dst)
}

func parseValues(values url.Values, dst interface{}) error {
	dec := schema.NewDecoder()
	dec.IgnoreUnknownKeys(true)

	return dec.Decode(dst, values)
}

func getUIntFromPath(r *http.Request, variable string) (uint, error) {
	res, err := strconv.Atoi(chi.URLParam(r, variable))
	if err != nil {
		return 0, err
	}

	return uint(res), nil
}
