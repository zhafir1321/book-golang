package main

import (
	"book-golang/configs"
	"book-golang/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	configs.ConnectDB()

	r := mux.NewRouter()

	router := r.PathPrefix("/api").Subrouter()

	routes.AuthRoutes(router)
	routes.BookRoutes(router)
	routes.BorrowRoutes(router)

	log.Println("Server started at :8080")
	http.ListenAndServe(":5000", router)
}