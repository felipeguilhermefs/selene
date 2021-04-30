package controllers

import (
	"log"
	"net/http"

	"github.com/felipeguilhermefs/selene/services"
	"github.com/felipeguilhermefs/selene/view"
)

// LoginController controls login endpoints
type LoginController struct {
	page     *view.View
	userSrvc services.UserService
}

// newLoginController creates a new instance of LoginController
func newLoginController(userSrvc services.UserService) *LoginController {
	return &LoginController{
		page:     view.NewView("login"),
		userSrvc: userSrvc,
	}
}

// loginForm data necessary to login a user
type loginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func (uc *LoginController) LoginPage(w http.ResponseWriter, r *http.Request) {
	var form loginForm
	parseURLParams(r, &form)
	uc.page.Render(w, r, view.NewData(&form))
}

func (uc *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	var form loginForm
	vd := view.NewData(&form)
	if err := parseForm(r, &form); err != nil {
		log.Println(err)
		uc.page.Render(w, r, vd.WithError(err))
		return
	}

	_, err := uc.userSrvc.Authenticate(form.Email, form.Password)
	if err != nil {
		uc.page.Render(w, r, vd.WithError(err))
		return
	}

	http.Redirect(w, r, "/books", http.StatusFound)
}
