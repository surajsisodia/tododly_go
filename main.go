package main

import (
	"fmt"
	"log"
	"net/http"
	"tododly/db"
	"tododly/middlewares"
	"tododly/routes"

	"github.com/gorilla/mux"
)

func main() {
	db.GetSqlDbConnection()

	router := mux.NewRouter()

	//router.Use(middlewares.AuthMiddleware)
	routes.TaskRoutes(router)
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.Use(middlewares.LoggerMiddleware)

	router.HandleFunc("/", serveHome).Methods("GET")

	fmt.Println("Server is up and running")
	log.Fatal(http.ListenAndServe(":8000", router))

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World</h1>"))
}
