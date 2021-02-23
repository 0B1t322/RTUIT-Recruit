package shop

import (
	"github.com/0B1t322/RTUIT-Recruit/pkg/models/product"
	"gorm.io/gorm"
)

// Shop a model of  shop
//
// If  you want to add current product just input theirs IDs  in Product
//
// If  want  to create some  new product input only product  data waithout ID
type Shop struct {
	ID			uint				`gorm:"primarykey"`
	Name		string 				`json:"name" gorm:"unique"`
	Adress		string				`json:"adress"`
	PhoneNubmer	string				`json:"phone_number"`

	// here all existing  products
	Products	[]product.Product	`json:"products"  gorm:"many2many:shop_products;"`
	// here product count
	// Counts		[]count.Count		`json:"counts" gorm:"-"`
}


func (s Shop) BeforeDelete(tx *gorm.DB)  error {
	var sp []ShopProducts
	if err := tx.Table("shop_products").Where("shop_id = ?", s.ID).Find(&sp).Error; err != nil {
		return err
	}

	if err := tx.Delete(sp).Error; len(sp) > 0 && err != nil {
		return err
	}

	return nil
}

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&Shop{}); err != nil {
		return err
	}

	if err :=  db.AutoMigrate(&ShopProducts{}); err != nil {
		return err
	}

	return nil
}