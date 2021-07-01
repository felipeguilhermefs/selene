package handlers

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/infrastructure/auth"
	"github.com/felipeguilhermefs/selene/view"
)

// loginForm data necessary to login a user
type loginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func HandleLoginPage(loginView *view.View) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var form loginForm
		parseURLParams(r, &form)
		loginView.Render(w, r, view.NewData(&form))
	}
}

func HandleLogin(loginView *view.View, authService auth.AuthService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var form loginForm
		vd := view.NewData(&form)

		err := parseForm(r, &form)
		if err != nil {
			loginView.Render(w, r, vd.WithError(err))
			return
		}

		err = authService.Login(w, r, form.Email, form.Password)
		if err != nil {
			loginView.Render(w, r, vd.WithError(err))
			return
		}

		http.Redirect(w, r, "/books", http.StatusFound)
	}
}
