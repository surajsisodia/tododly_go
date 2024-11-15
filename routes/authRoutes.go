package routes

import (
	"tododly/controllers"

	"github.com/gorilla/mux"
)

func AuthRoutes(router *mux.Router) *mux.Router {

	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
	router.HandleFunc("/api/signup", controllers.Signup).Methods("POST")

	return router
}
