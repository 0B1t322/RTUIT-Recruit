package product

import (
	model "github.com/0B1t322/RTUIT-Recruit/pkg/models/product"
	"gorm.io/gorm"
)

type ProductController struct {
	db *gorm.DB
}

func New(db *gorm.DB) *ProductController {
	return &ProductController{db: db}
}

func (pc *ProductController) Get(ID uint) (*model.Product, error) {
	p := &model.Product{}

	if err := pc.db.First(p, "ID = ?", ID).Error; err == gorm.ErrRecordNotFound {
		return nil,  ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return p, nil
}

func (pc *ProductController) Create(p *model.Product) error {
	if err := pc.db.Create(p).Error; err != nil {
		return err
	}

	return nil
}

func (pc *ProductController) Update(p *model.Product) error {
	if err := pc.db.Updates(p).Error; err != nil {
		return err
	}

	return nil
}

func (pc *ProductController) Delete(p *model.Product) error {
	if err := pc.db.Delete(p).Error; err != nil {
		return err
	}

	return nil
}