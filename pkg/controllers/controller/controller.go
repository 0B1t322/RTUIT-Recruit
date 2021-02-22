package controller

import "gorm.io/gorm"

type Controller struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Controller {
	return &Controller{db}
}

func (c *Controller) Get(ID uint) (interface{},  error) {
	var i interface{}

	if err := c.db.First(i, "ID = ?", ID).Error; err != nil {
		return nil, err
	}

	return i, nil
}

func (c *Controller) Create(i interface{}) error {
	if err := c.db.Create(i).Error; err != nil {
		return err
	}

	return nil
}

func (c *Controller) Update(i interface{}) error {
	if err := c.db.Updates(i).Error; err != nil {
		return err
	}

	return nil
}

func (c *Controller) Delete(i interface{}) error {
	if err := c.db.Delete(i).Error; err != nil {
		return err
	}

	return nil
}