package handlers

import (
	"net/http"

	"github.com/felipeguilhermefs/selene/core/auth"
	"github.com/felipeguilhermefs/selene/infrastructure/session"
	"github.com/felipeguilhermefs/selene/view"
)

// signupForm data necessary to create a user
type signupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
	Name     string `schema:"name"`
}

func HandleSignupPage(signupView *view.View) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var form signupForm
		parseURLParams(r, &form)
		signupView.Render(w, r, view.NewData(&form))
	}
}

func HandleSignup(signupView *view.View, userAdder auth.UserAdder, sessionStore session.SessionStore) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var form signupForm
		vd := view.NewData(&form)

		err := parseForm(r, &form)
		if err != nil {
			signupView.Render(w, r, vd.WithError(err))
			return
		}

		newUser := &auth.NewUser{
			Name:     form.Name,
			Email:    form.Email,
			Password: form.Password,
		}

		err = userAdder.Add(newUser)
		if err != nil {
			signupView.Render(w, r, vd.WithError(err))
			return
		}

		err = sessionStore.SignIn(w, r, newUser.Email)
		if err != nil {
			signupView.Render(w, r, vd.WithError(err))
			return
		}

		http.Redirect(w, r, "/books", http.StatusFound)
	}
}
