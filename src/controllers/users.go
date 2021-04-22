package controllers

import (
	"log"
	"net/http"

	"github.com/felipeguilhermefs/selene/models"
	"github.com/felipeguilhermefs/selene/services"
	"github.com/felipeguilhermefs/selene/view"
)

// UserController controls all user view and endpoints
type UserController struct {
	signupView *view.View
	userSrvc   services.UserService
}

// newUserController creates a new instance of UserController
func newUserController(userSrvc services.UserService) *UserController {
	return &UserController{
		signupView: view.NewView("signup"),
		userSrvc:   userSrvc,
	}
}

// signupForm data necessary to create a user
type signupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
	Name     string `schema:"name"`
}

// SignupPage GET /signup
func (uc *UserController) SignupPage(w http.ResponseWriter, r *http.Request) {
	var form signupForm
	parseURLParams(r, &form)
	uc.signupView.Render(w, r, view.NewData(&form))
}

// Signup POST /signup
func (uc *UserController) Signup(w http.ResponseWriter, r *http.Request) {
	var form signupForm
	vd := view.NewData(&form)
	if err := parseForm(r, &form); err != nil {
		log.Println(err)
		uc.signupView.Render(w, r, vd.WithError(err))
		return
	}

	user := &models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}
	if err := uc.userSrvc.Create(user); err != nil {
		log.Println(err)
		uc.signupView.Render(w, r, vd.WithError(err))
		return
	}

	http.Redirect(w, r, "/books", http.StatusFound)
}

