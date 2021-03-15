package router

import (
	h "github.com/0B1t322/RTUIT-Recruit/service.shops/shophandler"
	"net/http"

	"github.com/0B1t322/RTUIT-Recruit/pkg/middlewares"
	"github.com/gorilla/mux"
)

func New(h h.ShopHandler) *mux.Router {
	r := mux.NewRouter()
	// Get Shop
	r.HandleFunc("/shops/{id:[0-9]+}", h.Get).Methods("GET")

	// Buy product
	r.Handle(
		"/shops/{id:[0-9]+}/{pid:[0-9]+}", 
			http.HandlerFunc(h.Buy),
	).Methods("PUT")

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
	
	// Create new shop
	r.Handle(
			"/shops/", 
			middlewares.CheckTokenIfFromService(
					http.HandlerFunc(
						h.CreateShop,
					),
			),
	).Methods("POST")
	
	return r
}