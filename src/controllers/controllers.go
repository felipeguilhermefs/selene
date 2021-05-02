package controllers

import (
	"github.com/gorilla/mux"

	"github.com/felipeguilhermefs/selene/services"
)

// NewControllers init all Controllers
func NewControllers(router *mux.Router, srvcs *services.Services) *Controllers {

	loginCtrl := newLoginController(srvcs.User, srvcs.Session)
	signupCtrl := newSignupController(srvcs.User, srvcs.Session)

	router.HandleFunc("/login", loginCtrl.Login).Methods("POST")
	router.HandleFunc("/login", loginCtrl.LoginPage).Methods("GET")
	router.HandleFunc("/signup", signupCtrl.SignupPage).Methods("GET")
	router.HandleFunc("/signup", signupCtrl.Signup).Methods("POST")

	return &Controllers{
		login:  loginCtrl,
		signup: signupCtrl,
	}
}

// Controllers holds reference to all controllers
type Controllers struct {
	login  *LoginController
	signup *SignupController
}
