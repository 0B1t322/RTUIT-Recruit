package count

import "gorm.io/gorm"

type Count struct {
	// ID of current product
	ProductID	string 	`json:"product_id"`

	// ID of shop where this product exist
	ShopID		string	`json:"shop_id"`

	// Count of product
	Count		int		`json:"count"`
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Count{})
}