package purchase

import (
	"time"
	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model
	UID			string		
	BuyDate		time.Time 	`json:"buy_date"`
	ProductName	string		`json:"product_name"`
	Cost		float64		`json:"cost"`
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Purchase{})
}