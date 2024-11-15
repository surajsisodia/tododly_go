package routes

import (
	"net/http"
	"tododly/controllers"
	"tododly/middlewares"

	"github.com/gorilla/mux"
)

func TaskRoutes(router *mux.Router) {

	//router := mux.NewRouter()

	taskSub := router.Methods(http.MethodGet, http.MethodPost).Subrouter()

	taskSub.HandleFunc("/api/tasks", controllers.GetAllTasks).Methods("GET")
	taskSub.HandleFunc("/api/task", controllers.GetSingleTask).Methods("GET")
	taskSub.HandleFunc("/api/task", controllers.CreateNewTask).Methods("POST")

	taskSub.Use(middlewares.AuthMiddleware)

	// router.HandleFunc("/api/tasks", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetAllTasks))).Methods("GET")
	// router.HandleFunc("/api/task", controllers.GetSingleTask).Methods("GET")
	// router.HandleFunc("/api/task", controllers.CreateNewTask).Methods("POST")

	//return taskSub
}
