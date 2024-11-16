package routes

import (
	"book-golang/controllers"
	"book-golang/middlewares"

	"github.com/gorilla/mux"
)

func BorrowRoutes(r *mux.Router) {
	router := r.PathPrefix("/borrow").Subrouter()

	router.Use(middlewares.Authentication)
	router.HandleFunc("/{bookId}", controllers.CreateBorrow).Methods("POST")
	router.HandleFunc("/{bookId}", controllers.ReturnBook).Methods("PUT")
	router.HandleFunc("/", controllers.GetBorrows).Methods("GET")

}