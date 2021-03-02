package main

import (
	"flag"

	"github.com/0B1t322/RTUIT-Recruit/pkg/db"
	"github.com/0B1t322/RTUIT-Recruit/pkg/middlewares"
	"github.com/0B1t322/RTUIT-Recruit/service.purchases/app"
	log "github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()

	DB, err := db.DBManager.OpenDataBase(db.DBName)
	if err != nil {
		log.WithFields(log.Fields{
			"Package": "main",
			"Err": err,
		}).Panic()
	}
	
	a := app.New(DB, "8081")

	log.Info("Server started on :8081")
	log.Info("Key is: " + middlewares.SecretKey)
	if err := a.Start(); err != nil {
		panic(err)
	}	
}