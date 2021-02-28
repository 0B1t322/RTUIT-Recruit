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
	Shop		shop.Shop		`json:"shop" gorm:"foreignKey:ID;references:ShopID"`

	BuyDate		time.Time 		`json:"buy_date"`

	ProductID	uint			`json:"product_id"`
	Product		product.Product	`json:"product" gorm:"foreignKey:ID;references:ProductID"`

	// Payment can be cash/card
	Payment		string			`json:"payment"`

	Count		uint			`json:"count"`

	Cost		uint			`json:"cost"`
}

type jsonPurchase struct {
	ID			uint			`json:"id"`

	UID			uint			`json:"uid"`
	ShopID		uint			`json:"shop_id"`
	Shop		shop.ShopInfo	`json:"shop"`

	BuyDate		time.Time 		`json:"buy_date"`

	ProductID	uint			`json:"product_id"`
	Product		product.Product	`json:"product"`

	// Payment can be cash/card
	Payment		string			`json:"payment"`

	Count		uint			`json:"count"`

	Cost		uint			`json:"cost"`
}

func (p *Purchase) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonPurchase{
		ID: p.ID,
		UID: p.UID,
		ShopID: p.ShopID,
		Shop: p.Shop.ShopInfo,
		BuyDate: p.BuyDate,
		ProductID: p.ProductID,
		Product: p.Product,
		Payment: p.Payment,
		Count: p.Count,
		Cost: p.Cost,
	})
}

func (p *Purchase) UnmarshalJSON(data []byte) error {
	jsonPur := &jsonPurchase{}
	if err := json.Unmarshal(data, jsonPur); err != nil {
		return err
	}

	p.ID 		= jsonPur.ID
	p.UID		= jsonPur.UID
	p.ShopID	= jsonPur.ShopID
	p.Shop		= shop.Shop{ShopInfo: jsonPur.Shop}
	p.BuyDate	= jsonPur.BuyDate
	p.ProductID	= jsonPur.ProductID
	p.Product	= jsonPur.Product
	p.Payment	= jsonPur.Payment
	p.Count		= jsonPur.Count
	p.Cost		= jsonPur.Cost

	return nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Purchase{})
}

func (p *Purchase) AfterFind(tx *gorm.DB) error {

	if err := tx.Model(p).Association("Shop").Find(&p.Shop); err != nil {
		return err
	}

	if err := tx.Model(p).Association("Product").Find(&p.Product); err != nil {
		return err
	}

	return nil
}