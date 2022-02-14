package backend

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

	routes.UserRoutes(router)

	log.Fatal(http.ListenAndServe(":8000", router))
}
