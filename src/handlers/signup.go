package handlers

import (
	"log"
	"net/http"

	"github.com/felipeguilhermefs/selene/models"
	"github.com/felipeguilhermefs/selene/services"
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

func HandleSignup(signupView *view.View, authService services.AuthService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var form signupForm
		vd := view.NewData(&form)

		err := parseForm(r, &form)
		if err != nil {
			log.Println(err)
			signupView.Render(w, r, vd.WithError(err))
			return
		}

		newUser := models.User{
			Name:     form.Name,
			Email:    form.Email,
			Password: form.Password,
		}

		err = authService.SignUp(w, r, &newUser)
		if err != nil {
			signupView.Render(w, r, vd.WithError(err))
			return
		}

		http.Redirect(w, r, "/books", http.StatusFound)
	}
}
