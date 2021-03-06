package handlers

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/core/auth"
	"github.com/felipeguilhermefs/selene/infrastructure/session"
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

func HandleLogin(loginView *view.View, userVerifier auth.UserVerifier, sessionStore session.SessionStore) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var form loginForm
		vd := view.NewData(&form)

		err := parseForm(r, &form)
		if err != nil {
			loginView.Render(w, r, vd.WithError(err))
			return
		}

		user, err := userVerifier.Verify(form.Email, form.Password)
		if err != nil {
			loginView.Render(w, r, vd.WithError(err))
			return
		}

		err = sessionStore.SignIn(w, r, user.Email)
		if err != nil {
			loginView.Render(w, r, vd.WithError(err))
			return
		}

		http.Redirect(w, r, "/books", http.StatusFound)
	}
}
