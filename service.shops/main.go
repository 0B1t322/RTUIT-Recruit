package main

import (
	"github.com/0B1t322/RTUIT-Recruit/pkg/db"
	log "github.com/sirupsen/logrus"
	
	"github.com/0B1t322/RTUIT-Recruit/service.shops/app"
	"flag"
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

	a := app.New(DB, "8082", "service.purchases:8081")

	log.Infoln("Starting server on :8082")

	if err := a.Start(); err != nil {
		panic(err)
	}
}