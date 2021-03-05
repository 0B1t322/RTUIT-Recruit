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

func (sc *ShopController) GetAll() ([]m.Shop, error) {
	var shops []m.Shop

	if err := sc.db.Find(&shops).Error; err != nil {
		return nil, err
	}

	if len(shops) == 0 {
		return nil, ErrNotFound
	}

	return shops, nil

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

	return nil
}

// Update  update only shop basic  field
// 
// Don't update products and counts
func (sc *ShopController) Update(s *m.Shop) error {
	return sc.Controller.Update(s)
}

func (sc *ShopController) AddCount(ShopID, ProductID, Add uint) error {
	return addCount(
		sc.db,
		ShopID,
		ProductID,
		func(old uint) int {
			return int(old + Add)
		},
	)
}

func (sc *ShopController) SubCount(ShopID, ProductID, Sub uint) error {
	return addCount(
		sc.db,
		ShopID,
		ProductID,
		func(old uint) int {
			return int(old - Sub)
		},
	)
}

func addCount(
	db *gorm.DB,
	ShopID, ProductID uint,
	add func(old uint) int,
) error {
	// TODO think of how check if shop not exist
	// or product
	tx := db.Model(m.ShopProduct{}).
				Where(
					"shop_id = ? AND product_ID = ?",
					ShopID, ProductID,
				)
	var oldCount uint
	if err := tx.Select("count").First(&oldCount).Error; err == gorm.ErrRecordNotFound{
		return ErrProductNotFound
	} else if err != nil {
		return err
	}

	var newCount int
	if newCount = add(oldCount); newCount < 0 {
		return ErrNegCount
	}

	if err := tx.Update("count", newCount).Error; err != nil {
		return err
	}

	return nil
}