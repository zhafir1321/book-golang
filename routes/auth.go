package routes

import "github.com/gorilla/mux"

func AuthRoutes(r *mux.Router) {
	router := r.PathPrefix("/api").Subrouter()

}