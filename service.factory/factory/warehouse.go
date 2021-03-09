package factory

import (
	"sync"
)

type Warehouse struct {
	mu sync.Mutex
	prodcutsCount map[uint]uint
}

func NewWarehouse() *Warehouse {
	return &Warehouse{
		prodcutsCount: make(map[uint]uint),
	}
}

type Warehouser interface {
	AddProduct(productID, count uint)
	TakeAllProducts(productID uint) uint
}

func (w *Warehouse) AddProduct(productID, count uint) {
	w.mu.Lock()

	w.prodcutsCount[productID] += count

	w.mu.Unlock()
}

func (w *Warehouse) TakeAllProducts(productID uint) uint {
	w.mu.Lock()
	c, find := w.prodcutsCount[productID]
	if !find {
		return 0
	}

	defer func() {
		w.prodcutsCount[productID] = 0
		w.mu.Unlock()
	}()

	return c
}