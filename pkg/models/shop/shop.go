package shop

import (
	_ "github.com/0B1t322/RTUIT-Recruit/pkg/models/product"
	"gorm.io/gorm"
)

// Shop a model of  shop
//
// If  you want to add current product just input theirs IDs  in Product
//
// If  want  to create some  new product input only product  data waithout ID
type Shop struct {
	ShopInfo						

	// here product count
	ShopProducts		[]ShopProduct		`json:"shop_products" gorm:"foreignKey:ShopID;references:ID"`
}


func (s Shop) BeforeDelete(tx *gorm.DB)  error {
	if err := tx.Delete(s.ShopProducts).Error; len(s.ShopProducts) > 0 && err != nil {
		return err
	}

	return nil
}

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&Shop{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&ShopInfo{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&ShopProduct{}); err != nil {
		return err
	}

	return nil
}

func (s *Shop) setCount(tx *gorm.DB) error {
	if err := tx.Table("shop_products").
			Where("shop_id = ?", s.ID).
			Find(&s.ShopProducts).Error;
	err != nil {
		return err
	}

	return nil
}

func (s *Shop) AfterUpdate(tx *gorm.DB) error {
	for i, c := range s.ShopProducts {
		if err := tx.Table("shop_products").
					Where("shop_id = ? AND product_id = ? AND updated_at > ?", s.ID, c.ProductID, c.UpdatedAt).
					Find(&s.ShopProducts[i]).
					Error;
		err != nil {
			return err
		}
	}

	return nil
}

func (s *Shop) AfterFind(tx *gorm.DB) error {
	if err := tx.Find(&s.ShopProducts, "shop_id = ?", s.ID).Error; err != nil {
		return err
	}

	return nil
}