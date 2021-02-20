package db

import (
	log "github.com/sirupsen/logrus"
	"github.com/0B1t322/RTUIT-Recruit/pkg/models/purchase"
	"time"

	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
)

const DBName = "recruit?parseTime=true"

var DBManager *db.Manager

func init() {
	DBManager = db.NewManager("root", "root", "db:3306", 20*time.Second)

	DB, err := DBManager.OpenDataBase(DBName)
	if err != nil {
		panic(err)
	}

	// Mirgate models
	if err := purchase.AutoMigrate(DB); err != nil {
		log.Error(err)		
	}
}