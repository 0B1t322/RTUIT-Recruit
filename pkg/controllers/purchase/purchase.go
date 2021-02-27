package purchase

import (
	"strings"

	m "github.com/0B1t322/RTUIT-Recruit/pkg/models/purchase"
	"gorm.io/gorm"
)
// TODO base controller struct

type PurchaseController struct {
	db *gorm.DB
}

func New(db *gorm.DB) *PurchaseController {
	pc := &PurchaseController{db: db}

	return pc
}

func (pc *PurchaseController) Get(ID uint) (*m.Purchase, error) {
	p := &m.Purchase{}	
	if err := pc.db.First(p, "ID = ?", ID).Error; err == gorm.ErrRecordNotFound {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	// if err := pc.db.Model(p).Association("Shop").Find(&p.Shop); err != nil {
	// 	return nil, err
	// }

	// if err := pc.db.Model(p).Association("Product").Find(&p.Product); err != nil {
	// 	return nil, err
	// }

	return p, nil
}

func (pc *PurchaseController) GetAll(UID uint) ([]*m.Purchase, error)  {
	p := []*m.Purchase{}

	err := pc.db.Find(&p, "UID = ?", UID).Error

	if err != nil {
		return nil, err
	}

	if len(p) == 0 {
		return nil, ErrNotFound
	}

	// if err := pc.db.Model(p).Association("Shop").Find(&p.Shop); err != nil {
	// 	return nil, err
	// }

	// if err := pc.db.Model(p).Association("Product").Find(&p.Product); err != nil {
	// 	return nil, err
	// }

	return p, nil
}

func (pc *PurchaseController) Create(p *m.Purchase) error {
	if err := pc.db.Create(p).Error; err != nil && productNotFound(err) {
		return ErrInvalidProductID
	} else if err != nil && shopNotFound(err) {
		return ErrInvalidShopID
	} else if err != nil {
		return err
	}

	return nil
}

// TODO  maybe delete this because it's not to be updated
func (pc *PurchaseController) Update(p *m.Purchase) error {
	if err := pc.db.Updates(p).Error; err != nil {
		return err
	}

	return nil
}

func (pc *PurchaseController) Delete(p *m.Purchase) error {
	if err := pc.db.Delete(p).Error; err != nil {
		return err
	}

	return nil
}

// Два костыля

func shopNotFound(err error) bool {
	return strings.Contains(err.Error(), "`shops`")
}

func productNotFound(err error) bool {
	return strings.Contains(err.Error(), "`products`")
}