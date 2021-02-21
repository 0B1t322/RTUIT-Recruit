package shop

import "gorm.io/gorm"

type Shop struct {
	ID	uint	`gorm:"primarykey"`
	ShopInfo
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Shop{})
}