package routes

import (
	"net/http"
	"tododly/controllers"
	"tododly/middlewares"

	"github.com/gorilla/mux"
)

func TaskRoutes(router *mux.Router) {

	//router := mux.NewRouter()

	taskSub := router.PathPrefix("/api/task").Subrouter()
	// taskSub := router.Methods(http.MethodGet, http.MethodPost, http.MethodPatch).Subrouter()

	taskSub.HandleFunc("", controllers.GetAllTasks).Methods(http.MethodGet)
	taskSub.HandleFunc("/{task_id}", controllers.GetSingleTask).Methods(http.MethodGet)
	taskSub.HandleFunc("", controllers.CreateNewTask).Methods(http.MethodPost)
	taskSub.HandleFunc("/{task_id}", controllers.UpdateTask).Methods(http.MethodPatch)

	// privateRoute.Handler(LoggingMiddleware(AuthMiddleware(privateRoute.GetHandler())))
	taskSub.Use(middlewares.RequestBodyMiddleware)
	taskSub.Use(middlewares.AuthMiddleware)
	// taskSub.Use(middlewares.RequestBodyMiddleware)

	// router.HandleFunc("/api/tasks", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetAllTasks))).Methods("GET")
	// router.HandleFunc("/api/task", controllers.GetSingleTask).Methods("GET")
	// router.HandleFunc("/api/task", controllers.CreateNewTask).Methods("POST")

	//return taskSub
}
