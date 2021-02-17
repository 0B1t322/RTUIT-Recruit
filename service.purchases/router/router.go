package router

import (
	"github.com/0B1t322/RTUIT-Recruit/service.purchases/handlers"
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/purchases/{id:[0-9]+}", handlers.Get).Methods("GET")

	// r.HandleFunc("/purchases", handlers.GetAll).Methods("GET")

	r.HandleFunc("/purchases/", handlers.Add).Methods("POST")

	r.HandleFunc("/purchases/{id:[0-9]+}", handlers.Delete).Methods("DELETE")

	// r.HandleFunc("/purchases/{id:[0-9]+}", handlers.Update).Methods("PUT")

	return r
}
