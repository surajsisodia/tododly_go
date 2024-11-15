package routes

import (
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) *mux.Router {
	// router.HandleFunc("/api/user", controllers.GetAllTasks).Methods("GET")
	// router.HandleFunc("/api/signup", controllers.GetSingleTask).Methods("GET")

	return router
}
