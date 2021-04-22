package controllers

import (
	"github.com/gorilla/mux"

	"github.com/felipeguilhermefs/selene/services"
)

// NewControllers init all Controllers
func NewControllers(srvcs *services.Services) *Controllers {
	return &Controllers{
		user: newUserController(srvcs.User),
	}
}

// Controllers holds reference to all controllers
type Controllers struct {
	user *UserController
}

// RegisterRoutes
func (c *Controllers) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/signup", c.user.SignupPage).Methods("GET")
	router.HandleFunc("/signup", c.user.Signup).Methods("POST")
}
