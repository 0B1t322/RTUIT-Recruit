package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/0B1t322/RTUIT-Recruit/service.purchases/router"
)

func main() {
	log.Info("Server started on :8081")
	if err :=  http.ListenAndServe(":8081", router.New()); err  != nil {
		panic(err)
	}	
}