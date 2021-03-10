package factory

import (
	log "github.com/sirupsen/logrus"
	"sync"
)

type Warehouse struct {
	prodcutsCount sync.Map
	// mu sync.RWMutex
}

func NewWarehouse() *Warehouse {
	return &Warehouse{
		prodcutsCount: sync.Map{},
	}
}

type Warehouser interface {
	AddProduct(productID, count uint)
	TakeProduct(productID uint ) <- chan uint8
}

func (w *Warehouse) AddProduct(productID, count uint) {
	log.Infof("Start add prodct PID: %v count: %v", productID, count)
	c, find := w.prodcutsCount.Load(productID)
	if !find {
		w.prodcutsCount.Store(productID, make(chan uint8)) 
		c, _ = w.prodcutsCount.Load(productID)
	}

	cuint := c.(chan uint8)
	for i := uint(0); i < count; i++ {
		cuint <- 1
	}
	log.Infof("Product added")
}

func (w *Warehouse) TakeProduct( productID uint ) <- chan uint8 {
	log.Infof("Start take product %v", productID)
	c := w.getChanelWhenCreated(productID)
	log.Infof("return taked products")
	return c
}

func (w *Warehouse) getChanelWhenCreated(productID uint) chan uint8 {
	var c chan uint8
	created := false
	for !created {
		if _c, find := w.prodcutsCount.Load(productID); find {
			c = _c.(chan uint8)
			created = true
		}
	}

	return c
}