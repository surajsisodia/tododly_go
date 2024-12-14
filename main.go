package main

import (
	"fmt"
	"log"
	"net/http"
	"tododly/db"
	"tododly/middlewares"
	"tododly/routes"
	"tododly/utils"

	"github.com/gorilla/mux"
)

func main() {
	utils.LoadEnvVars()
	db.Connections = db.GetSqlDbConnection()

	if db.Connections == nil {
		log.Fatal("Database connection Failure")
		return
	}

	router := mux.NewRouter()

	//router.Use(middlewares.AuthMiddleware)
	// router.Use(middlewares.RequestBodyMiddleware)

	routes.TaskRoutes(router)
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.Use(middlewares.LoggerMiddleware)

	router.HandleFunc("/", serveHome).Methods("GET")

	fmt.Println("Server is up and running")
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World</h1>"))
}
