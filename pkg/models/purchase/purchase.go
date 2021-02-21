package purchase

import (
	"time"

	"github.com/0B1t322/RTUIT-Recruit/pkg/models/product"
	"github.com/0B1t322/RTUIT-Recruit/pkg/models/shop"
	"gorm.io/gorm"
)

type Purchase struct {
	ID			uint			`gorm:"primarykey"`

	UID			uint

	ShopID		uint			`json:"-"`
	Shop		shop.Shop		`json:"shop" gorm:"foreignKey:ShopID"`

	BuyDate		time.Time 		`json:"buy_date"`

	ProductID	uint			`json:"-"`
	Product		product.Product	`json:"product" gorm:"foreignKey:ProductID"`

	// Payment can be cash/card
	Payment		string			`json:"payment"`
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Purchase{})
}