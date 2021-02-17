package db

import (
	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	if _DB, err := db.DBManger.OpenDataBase("recruit?parseTime=true"); err != nil {
		panic(err)
	} else {
		DB = _DB
	}
}