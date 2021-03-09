package factory

import (
	"sync"

	"github.com/0B1t322/RTUIT-Recruit/service.factory/hash"

	"encoding/json"
	"fmt"
	"net/http"

	m "github.com/0B1t322/RTUIT-Recruit/pkg/models/shop"

	"github.com/0B1t322/RTUIT-Recruit/pkg/models/product"
	log "github.com/sirupsen/logrus"
)

var count chan uint

func init() {
	count = make(chan uint)
}

type ProductCreater interface {
	Create(*product.Product) error
}


type Factory struct {
	ProductionCapacity 	string					`json:"production_capacity"`
	ProductPerCap		uint					`json:"product_per_cap"`
	Products			[]product.Product		`json:"products"`
	ShopNetwork			string
	productsID			map[uint]hash.HashUint
	mu sync.Mutex
}

func logError(m string, err error) {
	log.WithFields(log.Fields{
		"method": m,
		"err": err,
	}).Error()
}

func (f *Factory) checkShops() {
	url := fmt.Sprintf("%s/shops", f.ShopNetwork)

	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logError("checkShop", err)
		return
	}

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		logError("checkShop", err)
		return
	}
	defer resp.Body.Close()

	var shops []m.Shop
	{
		d := json.NewDecoder(resp.Body)
		if err := d.Decode(&shops); err != nil {
			logError("checkShops", err)
			return
		}
	}

	f.mu.Lock()
	defer f.mu.Unlock()

	for _, shop := range shops {
		for _, shopProduct := range shop.ShopProducts {
			h, find := f.productsID[shopProduct.ProductID]
			if !find {
				f.productsID[shopProduct.ProductID] = make(hash.HashUint)
				h = f.productsID[shopProduct.ProductID]
			}

			h.Add(shopProduct.ShopID)
		}
	}

}

func (f *Factory) createProducts(pc ProductCreater) {
	for _, p := range f.Products {
		if err := pc.Create(&p); err != nil {
			logError("createProducts", err)
		}
	}
}

func (f *Factory) makeProducts(w Warehouser) {
	c := f.ProductPerCap
	for productID := range f.productsID {
		go w.AddProduct(productID, c)
	}
}

func (f *Factory) sendToDelivery() {
	
}




// TODO сделать очередь порядка доставки в магазины

// func (f *Factory) makeProduct(PID uint) {
// 	shopID := f.productsID[PID]
// 	r, err := http.NewRequest("PUT", fmt.Sprintf("%s/shops/%v/%v/%v", f.ShopNetwork, shopID, PID, f.ProductPerCap), nil)
// 	if err != nil {
// 		logError("makeProduct", err)
// 		return
// 	}

// 	resp, err := http.DefaultClient.Do(r)
// 	if err != nil {
// 		logError("makeProduct", err)
// 		return
// 	}

// 	if resp.StatusCode == http.StatusBadRequest {

// 	}
// }