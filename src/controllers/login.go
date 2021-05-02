package controllers

import (
	"log"
	"net/http"

	"github.com/felipeguilhermefs/selene/services"
	"github.com/felipeguilhermefs/selene/view"
)

// LoginController controls login endpoints
type LoginController struct {
	page        *view.View
	userSrvc    services.UserService
	sessionSrvc services.SessionService
}

// newLoginController creates a new instance of LoginController
func newLoginController(
	userSrvc services.UserService,
	sessionSrvc services.SessionService,
) *LoginController {
	return &LoginController{
		page:        view.NewView("login"),
		sessionSrvc: sessionSrvc,
		userSrvc:    userSrvc,
	}
}

// loginForm data necessary to login a user
type loginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func (lc *LoginController) LoginPage(w http.ResponseWriter, r *http.Request) {
	var form loginForm
	parseURLParams(r, &form)
	lc.page.Render(w, r, view.NewData(&form))
}

func (lc *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	var form loginForm
	vd := view.NewData(&form)
	if err := parseForm(r, &form); err != nil {
		log.Println(err)
		lc.page.Render(w, r, vd.WithError(err))
		return
	}

	user, err := lc.userSrvc.Authenticate(form.Email, form.Password)
	if err != nil {
		lc.page.Render(w, r, vd.WithError(err))
		return
	}

	err = lc.sessionSrvc.SignIn(w, r, user)
	if err != nil {
		lc.page.Render(w, r, vd.WithError(err))
		return
	}

	http.Redirect(w, r, "/books", http.StatusFound)
}
