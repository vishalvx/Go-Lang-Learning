package router

import (
	"github.com/gorilla/mux"
	"github.com/vishalvx/dbconnect/controller"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	// Handle Home Route
	router.HandleFunc("/", controller.GetHome).Methods("GET")
	// Handles Movies Endpoints
	router.HandleFunc("/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/movies", controller.CreateOneMovie).Methods("POST")
	router.HandleFunc("/movies/watch/{id}", controller.MarkedAsWatched).Methods("PUT")
	router.HandleFunc("/movies/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/movies", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
