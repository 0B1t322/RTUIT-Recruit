package shop

import (
	"gorm.io/gorm"
)

type ShopProducts struct {
	ShopID		uint	`gorm:"primaryKey"`
	ProductID	uint	`gorm:"primaryKey"`
	Count		uint
}

func (ShopProducts) BeforeCreate(db *gorm.DB) error {
	return nil
}

func SetupJoinTable(db *gorm.DB) error {
	return db.SetupJoinTable(&Shop{}, "Products", &ShopProducts{})
}