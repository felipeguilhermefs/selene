package controllers

import (
	"github.com/gorilla/mux"

	"github.com/felipeguilhermefs/selene/services"
)

// NewControllers init all Controllers
func NewControllers(router *mux.Router, srvcs *services.Services) *Controllers {

	loginCtrl := newLoginController(srvcs.User, srvcs.Session)

	router.HandleFunc("/login", loginCtrl.Login).Methods("POST")
	router.HandleFunc("/login", loginCtrl.LoginPage).Methods("GET")

	return &Controllers{
		login:  loginCtrl,
	}
}

// Controllers holds reference to all controllers
type Controllers struct {
	login  *LoginController
}
