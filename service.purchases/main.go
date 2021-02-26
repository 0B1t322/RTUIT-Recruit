package main

import (
	"flag"
	"github.com/0B1t322/RTUIT-Recruit/service.purchases/app"
	"github.com/0B1t322/RTUIT-Recruit/pkg/db"
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
	if err := a.Start(); err != nil {
		panic(err)
	}	
}