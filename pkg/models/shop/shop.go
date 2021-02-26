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
	ShopInfo

	// here all existing  products
	Products	[]product.Product	`json:"products"  gorm:"many2many:shop_products;"`
	// here product count
	Count		[]ShopProducts		`json:"-" gorm:"-"`
}


func (s Shop) BeforeDelete(tx *gorm.DB)  error {
	if err := tx.Delete(s.Count).Error; len(s.Count) > 0 && err != nil {
		return err
	}

	return nil
}

func (s *Shop) AfterCreate(tx *gorm.DB) error {
	for i, p := range s.Products {
		if p.ID != 0 {
			err := tx.First(&s.Products[i]).Error
			if err != nil {
				return err
			}
		}
	}

	if err := s.setCount(tx); err != nil {
		return err
	}

	return nil
}

func (s *Shop) AfterFind(tx *gorm.DB) error {
	return s.setProductAndCounts(tx)
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

func (s *Shop) setProducts(tx *gorm.DB) error {
	var ProductsID []uint
	if err := tx.Table("shop_products").
				Where("shop_id = ?", s.ID).
				Select("product_id").
				Find(&ProductsID).Error; err != nil {
		return err
	}

	if err := tx.Model(s.Products).
				Where(ProductsID).
				Find(&s.Products).
				Error; 
	err != nil {
		return err
	}

	return nil
}

func (s *Shop) setCount(tx *gorm.DB) error {
	if err := tx.Table("shop_products").
			Where("shop_id = ?", s.ID).
			Find(&s.Count).Error;
	err != nil {
		return err
	}

	return nil
}

func (s *Shop) setProductAndCounts(tx *gorm.DB) error {
	if err := tx.Model(s.Count).
			Where("shop_id = ?", s.ID).
			Find(&s.Count).Error;
	err != nil {
		return err
	}

	var ProductsID []uint
	for _, c := range s.Count {
		ProductsID = append(ProductsID, c.ProductID)
	}

	if err := tx.Model(s.Products).
				Where(ProductsID).
				Find(&s.Products).
				Error;
	err != nil {
		return err
	}

	return nil
}

func (s *Shop) AfterUpdate(tx *gorm.DB) error {
	for i, c := range s.Count {
		if err := tx.Table("shop_products").
					Where("shop_id = ? AND product_id = ? AND updated_at > ?", s.ID, c.ProductID, c.UpdatedAt).
					Find(&s.Count[i]).
					Error;
		err != nil {
			return err
		}
	}

	return nil
}