package router

import (
	"net/http"

	"github.com/0B1t322/RTUIT-Recruit/pkg/middlewares"
	"github.com/0B1t322/RTUIT-Recruit/service.shops/handlers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()
	h := handlers.New(db)
	h.PurhacesNetwork = "http://localhost:8081"
	// Get Shop
	r.HandleFunc("/shops/{id:[0-9]+}", h.Get).Methods("GET")

	// Buy product
	r.HandleFunc("/shops/{id:[0-9]+}/{pid:[0-9]+}", h.Buy).Methods("PUT")

	// Check purchases
	r.HandleFunc("/shops/purchases/{uid:[0-9]+}", h.GetPurchases).Methods("GET")

	// Add count product
	// this method only for fabrics
	r.Handle(
			"/shops/{id:[0-9]+}/{pid:[0-9]+}/{count:[0-9]+}", 
			middlewares.CheckTokenIfFromService(
				http.HandlerFunc(h.AddCount),
			),
		).Methods("PUT")

	// Add product to shop
	r.Handle(
			"/shops/{id:[0-9]+}/{pid:[0-9]+}", 
			middlewares.CheckTokenIfFromService(
				http.HandlerFunc(h.AddProduct),
			),
		).Methods("POST")

	// Get All Shop
	r.HandleFunc("/shops", h.GetAll).Methods("GET")

	return r
}