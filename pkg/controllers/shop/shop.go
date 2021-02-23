package shop

import (
	m "github.com/0B1t322/RTUIT-Recruit/pkg/models/shop"
	"github.com/0B1t322/RTUIT-Recruit/pkg/controllers/controller"
	"gorm.io/gorm"
)

type ShopController struct {
	db *gorm.DB
	*controller.Controller
}

func New(db *gorm.DB) *ShopController {
	return  &ShopController{
		db:  db,
		Controller: controller.New(db),
	}
}

func (sc *ShopController) Get(ID uint) (*m.Shop, error) {
	s := &m.Shop{}
	
	if err := sc.db.First(s, "id = ?", ID).Error; err == gorm.ErrRecordNotFound {
		return nil, ErrNotFound
	} else if err !=  nil {
		return nil,  err
	}

	return s, nil
}

func (sc *ShopController) Delete(s *m.Shop) error {
	if err := sc.Controller.Delete(s); err != nil {
		return err
	}

	return nil
}

func (sc *ShopController) Create(s *m.Shop) error {
	if err := sc.Controller.Create(s); err != nil {
		return err
	}

	for i, p := range s.Products {
		if p.ID != 0 {
			defer func() {
				sc.db.First(&s.Products[i])
			}()
		}
	}

	return nil
}

// Update  update only shop basic  field
// 
// Don't update products and counts
func (sc *ShopController) Update(s *m.Shop) error {
	return  sc.Controller.Update(s)
}

// TODO methods to update count