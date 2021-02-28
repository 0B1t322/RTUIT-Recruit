package router

import (
	_ "github.com/0B1t322/RTUIT-Recruit/pkg/middlewares"
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter()

	// Get Shop
	r.HandleFunc("/shops/{id:[0-9]+}", nil).Methods("GET")

	// Buy product
	r.HandleFunc("/shops/{id:[0-9]+}/{pid:[0-9]+}", nil).Methods("PUT")

	// Check purchases
	r.HandleFunc("/shops/{uid:[0-9]+}", nil).Methods("GET")

	// Add product
	// this method only for fabrics
	r.Handle("/shops/{id:[0-9]+}/{pid:[0-9]+}", nil).Methods("POST")

	// Get All Shop
	r.HandleFunc("/shops", nil).Methods("GET")

	return r
}