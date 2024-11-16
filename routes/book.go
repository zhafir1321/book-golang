package routes

import (
	"book-golang/controllers"
	"book-golang/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func BookRoutes(r *mux.Router) {
	router := r.PathPrefix("/book").Subrouter()

	router.Use(middlewares.Authentication)
	router.HandleFunc("/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/{id}", controllers.GetBook).Methods("GET")
	router.Handle("/", middlewares.Authorization(http.HandlerFunc(controllers.CreateBook))).Methods("POST")
	router.HandleFunc("/{id}", controllers.UpdateBook).Methods("PUT")
	router.Handle("/{id}", middlewares.Authorization(http.HandlerFunc(controllers.DeleteBook))).Methods("DELETE")

}