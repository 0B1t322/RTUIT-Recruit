package purchase

import (
	m "github.com/0B1t322/RTUIT-Recruit/pkg/models/purchase"

	"github.com/0B1t322/RTUIT-Recruit/pkg/controller"
	"gorm.io/gorm"
)

type PurchaseController controller.Controller

func New(c *controller.Controller) *PurchaseController {
	pc := &PurchaseController{c.DB}

	return pc
}

func (pc *PurchaseController) Get(ID string) (*m.Purchase, error) {
	p := &m.Purchase{}	
	if err := pc.DB.First(p, "ID = ?", ID).Error; err == gorm.ErrRecordNotFound {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return p, nil
}

func (pc *PurchaseController) GetAll(UID string) ([]*m.Purchase, error)  {
	p := []*m.Purchase{}

	err := pc.DB.Find(&p, "UID = ?", UID).Error

	if err != nil {
		return nil, err
	}

	if len(p) == 0 {
		return nil, ErrNotFound
	}

	return p, nil
}

func (pc *PurchaseController) Create(p *m.Purchase) error {
	if err := pc.DB.Create(p).Error; err != nil {
		return err
	}

	return nil
}

func (pc *PurchaseController) Update(p *m.Purchase) error {
	if err := pc.DB.Updates(p).Error; err != nil {
		return err
	}

	return nil
}

func (pc *PurchaseController) Delete(p *m.Purchase) error {
	if err := pc.DB.Delete(p).Error; err != nil {
		return err
	}

	return nil
}