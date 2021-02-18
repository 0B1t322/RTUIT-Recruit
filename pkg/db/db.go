package db

import (
	log "github.com/sirupsen/logrus"
	"github.com/0B1t322/RTUIT-Recruit/pkg/models/purchase"
	"time"

	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
	"gorm.io/gorm"
)

const dbName = "recruit?parseTime=true"

var DB *gorm.DB

func init() {
	m := db.NewManager("root", "root", "db:3306", 20*time.Second)

	if _DB, err := m.OpenDataBase(dbName); err != nil {
		panic(err)
	} else {
		DB = _DB
	}

	// Mirgate models
	if err := purchase.AutoMigrate(DB); err != nil {
		log.Error(err)		
	}
}