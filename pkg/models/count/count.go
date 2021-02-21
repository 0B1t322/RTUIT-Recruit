package count

import "gorm.io/gorm"

type Count struct {
	// ID of current product
	ProductID	uint 	`json:"product_id"`

	// ID of shop where this product exist
	ShopID		uint	`json:"shop_id"`

	// Count of product
	Count		uint	`json:"count"`
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Count{})
}