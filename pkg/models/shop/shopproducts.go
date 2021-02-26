package shop

import (
	"time"

	"gorm.io/gorm"
)

type ShopProducts struct {
	ShopID		uint		`json:"shop_id" gorm:"primaryKey"`
	ProductID	uint		`json:"product_id" gorm:"primaryKey"`

	// TODO make foreign key here for product

	Count		uint		`json:"count"`
	UpdatedAt	time.Time
}

func (ShopProducts) BeforeCreate(db *gorm.DB) error {
	return nil
}

func SetupJoinTable(db *gorm.DB) error {
	return db.SetupJoinTable(&Shop{}, "Products", &ShopProducts{})
}