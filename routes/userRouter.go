package routes

import (
	"net/http"
	"tododly/controllers"
	"tododly/middlewares"

	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) *mux.Router {

	userSub := router.PathPrefix("/api/user").Subrouter()

	userSub.HandleFunc("", controllers.GetMyProfile).Methods(http.MethodGet)
	userSub.HandleFunc("", controllers.UpdateMyProfile).Methods(http.MethodPatch)

	userSub.Use(middlewares.AuthMiddleware, middlewares.RequestBodyMiddleware)

	return router
}
