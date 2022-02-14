package routes

import "github.com/gorilla/mux"

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/users")
	router.HandleFunc("/user/{id}")
	router.HandleFunc("/users")
	router.HandleFunc("/users/{id}")
	router.HandleFunc("/users/{id}")
}
