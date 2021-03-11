package router

import (
	"net/http"

	"github.com/0B1t322/RTUIT-Recruit/pkg/middlewares"
	ph "github.com/0B1t322/RTUIT-Recruit/service.purchases/purchaseshandler"
	"github.com/gorilla/mux"
)

//New return a new router
func New(ph ph.PurchaseHandler) *mux.Router {
	r := mux.NewRouter()

	r.Handle(
		"/purchases/{uid:[0-9]+}/{id:[0-9]+}", 
		middlewares.ContentTypeJSONMiddleware(
			http.HandlerFunc(
				ph.Get,
			),
		),
	).Methods("GET")

	r.Handle(
		"/purchases/{uid:[0-9]+}", 
		middlewares.ContentTypeJSONMiddleware(
			http.HandlerFunc(
				ph.GetAll,
			),
		),
	).Methods("GET")

	r.Handle(
		"/purchases/{uid:[0-9]+}", 
		middlewares.CheckTokenIfFromService(
			http.HandlerFunc(
				ph.Add,
			),
		),
	).Methods("POST")

	r.HandleFunc(
		"/purchases/{uid:[0-9]+}/{id:[0-9]+}", 
		ph.Delete,
	).Methods("DELETE")

	// r.HandleFunc("/purchases/{id:[0-9]+}", handlers.Update).Methods("PUT")

	// r.Use(middlewares.ContentTypeJSONMiddleware)

	return r
}
