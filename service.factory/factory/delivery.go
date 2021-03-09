package factory

import (
	"fmt"

	"github.com/0B1t322/RTUIT-Recruit/service.factory/hash"
)

type Delivery struct {
	ShopNetwork			string
}

type Deliverer interface {
	Send(m map[uint]hash.HashUint)
}

func (d *Delivery) Send(m map[uint]hash.HashUint) {
	var shopsID []uint
	for productID, h := range m {
		shopsID = h.GetKeys()
		for _, shopID := range shopsID {
			go d.send(productID, shopID)
		}
	}
}

func (d *Delivery) send(PID, SID uint) {
	url := fmt.Sprintf("%s/shops/%v/%v", d.ShopNetwork, SID, PID)
	
}