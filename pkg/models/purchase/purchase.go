package purchase

import (
	"encoding/json"
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
	ID			uint			`json:"id" gorm:"primarykey"`

	UID			uint			`json:"uid"`

	ShopID		uint			`json:"shop_id"`
	Shop		shop.Shop		`json:"shop" gorm:"foreignKey:ShopID"`

	BuyDate		time.Time 		`json:"buy_date"`

	ProductID	uint			`json:"product_id"`
	Product		product.Product	`json:"product" gorm:"foreignKey:ProductID"`

	// Payment can be cash/card
	Payment		string			`json:"payment"`

	Count		uint			`json:"count"`
}

type jsonPurchase struct {
	ID			uint			`json:"id"`

	UID			uint			`json:"uid"`

	Shop		shop.ShopInfo	`json:"shop"`

	BuyDate		time.Time 		`json:"buy_date"`

	Product		product.Product	`json:"product"`

	// Payment can be cash/card
	Payment		string			`json:"payment"`

	Count		uint			`json:"count"`
}

func (p *Purchase) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonPurchase{
		ID: p.ID,
		UID: p.UID,
		Shop: p.Shop.ShopInfo,
		BuyDate: p.BuyDate,
		Product: p.Product,
		Payment: p.Payment,
		Count: p.Count,
	})
}

func (p *Purchase) UnmarshalJSON(data []byte) error {
	jsonPur := &jsonPurchase{}
	if err := json.Unmarshal(data, jsonPur); err != nil {
		return err
	}

	p.ID 		= jsonPur.ID
	p.UID		= jsonPur.UID
	p.ShopID	= jsonPur.Shop.ID
	p.Shop		= shop.Shop{ShopInfo: jsonPur.Shop}
	p.BuyDate	= jsonPur.BuyDate
	p.ProductID	= jsonPur.Product.ID
	p.Product	= jsonPur.Product
	p.Payment	= jsonPur.Payment
	p.Count		= jsonPur.Count

	return nil
}
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Purchase{})
}