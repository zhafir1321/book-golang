package routes

import (
	"book-golang/controllers"

	"github.com/gorilla/mux"
)

func BookRoutes(r *mux.Router) {
	router := r.PathPrefix("/book").Subrouter()

	router.HandleFunc("/", controllers.GetBooks).Methods("GET")
	// router.HandleFunc("/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/", controllers.CreateBook).Methods("POST")
	// router.HandleFunc("/{id}", controllers.UpdateBook).Methods("PUT")
	// router.HandleFunc("/{id}", controllers.DeleteBook).Methods("DELETE")

}