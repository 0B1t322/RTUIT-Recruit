package shop

import (
	pc "github.com/0B1t322/RTUIT-Recruit/pkg/controllers/product"
	"time"

	"github.com/0B1t322/RTUIT-Recruit/pkg/models/product"
	"gorm.io/gorm"
)

type ShopProduct struct {
	ShopID		uint			`json:"shop_id" gorm:"primaryKey"`
	ProductID	uint			`json:"product_id" gorm:"primaryKey"`
	Product		product.Product	`json:"product" gorm:"foreignKey:ID;references:ProductID"`

	Count		uint			`json:"count"`

	UpdatedAt	time.Time
}

func (sp *ShopProduct) AfterCreate(tx *gorm.DB) error {
	return sp.findProduct(tx)
}

func (sp *ShopProduct) AfterFind(tx *gorm.DB) error {
	return sp.findProduct(tx)
}

func (sp *ShopProduct) findProduct(tx *gorm.DB) error {
	if sp.ProductID != 0 {
		err := tx.First(&sp.Product, "id = ?", sp.ProductID).Error
		if err == gorm.ErrRecordNotFound {
			return pc.ErrNotFound
		} else if  err !=  nil {
			return err
		}
	}

	return nil
}