package purchase

import (
	"github.com/0B1t322/RTUIT-Recruit/pkg/models/shop"
	"time"
	"gorm.io/gorm"
)

type Purchase struct {
	ID			uint			`gorm:"primarykey"`

	UID			uint

	ShopID		uint			`json:"-"`
	Shop		shop.Shop		`json:"shop" gorm:"foreignKey:ShopID"`

	BuyDate		time.Time 		`json:"buy_date"`
	ProductName	string			`json:"product_name"`
	Cost		float64			`json:"cost"`

	// Category can be set by a user or by a shop
	Category	string			`json:"category"`

	// Payment can be cash/card
	Payment		string			`json:"payment"`
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Purchase{})
}