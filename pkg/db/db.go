package db

import (
	log "github.com/sirupsen/logrus"
	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
	"gorm.io/gorm"
)

const dbName = "recruit?parseTime=true"

var DB *gorm.DB

func init() {
	log.Info(dbName)
	if _DB, err := db.DBManger.OpenDataBase(dbName); err != nil {
		panic(err)
	} else {
		DB = _DB
	}
}