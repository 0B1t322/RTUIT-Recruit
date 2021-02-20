package router

import (
	"github.com/0B1t322/RTUIT-Recruit/pkg/middlewares"
	"github.com/0B1t322/RTUIT-Recruit/service.purchases/handlers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

//New return a new router
func New(DB *gorm.DB) *mux.Router {
	r := mux.NewRouter()
	ph := handlers.New(DB)

	r.HandleFunc("/purchases/{uid:[0-9]+}/{id:[0-9]+}", ph.Get).Methods("GET")

	r.HandleFunc("/purchases/{uid:[0-9]+}", ph.GetAll).Methods("GET")

	r.HandleFunc("/purchases/{uid:[0-9]+}", ph.Add).Methods("POST")

	r.HandleFunc("/purchases/{uid:[0-9]+}/{id:[0-9]+}", ph.Delete).Methods("DELETE")

	// r.HandleFunc("/purchases/{id:[0-9]+}", handlers.Update).Methods("PUT")

	r.Use(middlewares.ContentTypeJSONMiddleware)

	return r
}
