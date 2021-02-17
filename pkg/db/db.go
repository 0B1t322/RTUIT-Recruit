package db

import (
	log "github.com/sirupsen/logrus"
	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
	"gorm.io/gorm"
)

const DBName = "recruit?parseTime=true"

var DB *gorm.DB

func init() {
	log.Info(DBName)
	if _DB, err := db.DBManger.OpenDataBase(DBName); err != nil {
		panic(err)
	} else {
		DB = _DB
	}
}