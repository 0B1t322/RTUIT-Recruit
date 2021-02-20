package purchase

import (
	m "github.com/0B1t322/RTUIT-Recruit/pkg/models/purchase"
	"gorm.io/gorm"
)

type PurchaseController struct {
	db *gorm.DB
}

func New(db *gorm.DB) *PurchaseController {
	pc := &PurchaseController{db: db}

	return pc
}

func (pc *PurchaseController) Get(ID string) (*m.Purchase, error) {
	p := &m.Purchase{}	
	if err := pc.db.First(p, "ID = ?", ID).Error; err == gorm.ErrRecordNotFound {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return p, nil
}

func (pc *PurchaseController) GetAll(UID string) ([]*m.Purchase, error)  {
	p := []*m.Purchase{}

	err := pc.db.Find(&p, "UID = ?", UID).Error

	if err != nil {
		return nil, err
	}

	if len(p) == 0 {
		return nil, ErrNotFound
	}

	return p, nil
}

func (pc *PurchaseController) Create(p *m.Purchase) error {
	if err := pc.db.Create(p).Error; err != nil {
		return err
	}

	return nil
}

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