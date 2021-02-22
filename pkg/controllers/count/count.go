package count

import (
	m "github.com/0B1t322/RTUIT-Recruit/pkg/models/count"
	"github.com/0B1t322/RTUIT-Recruit/pkg/controllers/controller"
	"gorm.io/gorm"
)

type CountController struct {
	db *gorm.DB
	*controller.Controller
}

func New(db *gorm.DB) *CountController{
	return &CountController{
		db: db,
		Controller: controller.New(db),
	}
}

func (cc *CountController) Get(ProductID, ShopID uint) (*m.Count, error) {
	c := &m.Count{}
	if err := cc.db.First(
		c, 
		"shop_id = ? AND product_id = ?", 
		ShopID, ProductID,
	).Error; err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		} else if err != nil {
			return nil, err
		}

	return c, nil
}

func (cc *CountController) Create(c *m.Count) error {
	if _, err := cc.Get(c.ProductID, c.ShopID); err != ErrNotFound {
		return ErrExist
	}

	return cc.Controller.Create(c)
}

func (cc *CountController) Delete(c *m.Count) error {
	if err := cc.db.Delete(
		c, 
		"shop_id = ? AND product_id = ?",
		c.ShopID, c.ProductID,
	).Error; err != nil {
		return err
	}

	return nil
}


func (cc *CountController) Update(c *m.Count) error {
	return cc.db.Model(c).Where(
		"shop_id = ? AND product_id = ?",
		c.ShopID, c.ProductID,
	).Update("count", c.Count).Error
}