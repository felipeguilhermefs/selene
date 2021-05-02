package controllers

import (
	"log"
	"net/http"

	"github.com/felipeguilhermefs/selene/models"
	"github.com/felipeguilhermefs/selene/services"
	"github.com/felipeguilhermefs/selene/view"
)

// SignupController controls singup and endpoints
type SignupController struct {
	page        *view.View
	userSrvc    services.UserService
	sessionSrvc services.SessionService
}

// newSignupController creates a new instance of SignupController
func newSignupController(
	userSrvc services.UserService,
	sessionSrvc services.SessionService,
) *SignupController {
	return &SignupController{
		page:        view.NewView("signup"),
		sessionSrvc: sessionSrvc,
		userSrvc:    userSrvc,
	}
}

// signupForm data necessary to create a user
type signupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
	Name     string `schema:"name"`
}

func (sc *SignupController) SignupPage(w http.ResponseWriter, r *http.Request) {
	var form signupForm
	parseURLParams(r, &form)
	sc.page.Render(w, r, view.NewData(&form))
}

func (sc *SignupController) Signup(w http.ResponseWriter, r *http.Request) {
	var form signupForm
	vd := view.NewData(&form)
	if err := parseForm(r, &form); err != nil {
		log.Println(err)
		sc.page.Render(w, r, vd.WithError(err))
		return
	}

	user := &models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}
	if err := sc.userSrvc.Create(user); err != nil {
		log.Println(err)
		sc.page.Render(w, r, vd.WithError(err))
		return
	}

	err := sc.sessionSrvc.SignIn(w, r, user)
	if err != nil {
		sc.page.Render(w, r, vd.WithError(err))
		return
	}

	http.Redirect(w, r, "/books", http.StatusFound)
}
