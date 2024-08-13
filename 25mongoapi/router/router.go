package router

import (
	"github.com/gorilla/mux"
	"github.com/rchethan602/mongoapi/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.GetAllMyMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteAmovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controller.DeleteAllmovies).Methods("DELETE")

	return router
}
