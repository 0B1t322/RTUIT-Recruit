package purchase

import (
	"time"

	"github.com/0B1t322/RTUIT-Recruit/pkg/models/product"
	"github.com/0B1t322/RTUIT-Recruit/pkg/models/shop"
	"gorm.io/gorm"
)
// Purchase is  a model  of  purchase
// if have shopID or ProductID  will search it in  DB
// if dont find them  returns errors
// 
// also you can create  this  shop or magazin by input  Product or magazin obj
type Purchase struct {
	ID			uint			`gorm:"primarykey"`

	UID			uint

	ShopID		uint			`json:"shop_id"`
	Shop		shop.Shop		`json:"shop" gorm:"foreignKey:ShopID"`

	BuyDate		time.Time 		`json:"buy_date"`

	ProductID	uint			`json:"product_id"`
	Product		product.Product	`json:"product" gorm:"foreignKey:ProductID"`

	// Payment can be cash/card
	Payment		string			`json:"payment"`

	Count		uint			`json:"count"`
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Purchase{})
}