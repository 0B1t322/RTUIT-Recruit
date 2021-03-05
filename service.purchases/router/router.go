package router

import (
	ph "github.com/0B1t322/RTUIT-Recruit/service.purchases/purchaseshandler"
	"github.com/0B1t322/RTUIT-Recruit/pkg/middlewares"
	"github.com/gorilla/mux"
)

//New return a new router
func New(ph ph.PurchaseHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/purchases/{uid:[0-9]+}/{id:[0-9]+}", ph.Get).Methods("GET")

	r.HandleFunc("/purchases/{uid:[0-9]+}", ph.GetAll).Methods("GET")

	r.HandleFunc("/purchases/{uid:[0-9]+}", ph.Add).Methods("POST")

	r.HandleFunc("/purchases/{uid:[0-9]+}/{id:[0-9]+}", ph.Delete).Methods("DELETE")

	// r.HandleFunc("/purchases/{id:[0-9]+}", handlers.Update).Methods("PUT")

	r.Use(middlewares.ContentTypeJSONMiddleware)
	r.Use(middlewares.CheckTokenIfFromService)

	return r
}
