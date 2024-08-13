package router

import (
	"bookstore/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	var r = mux.NewRouter()

	r.HandleFunc("/books", controller.GetAllBooks).Methods("GET")
	r.HandleFunc("/book", controller.InsertAbook).Methods("POST")
	return r
}
