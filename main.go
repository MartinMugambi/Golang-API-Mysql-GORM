package main

import (
	"backend/config"
	"backend/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	config.InitializeDatabase()

	routes.UserRouter(router)

	log.Fatal(http.ListenAndServe(":8000", router))
}
