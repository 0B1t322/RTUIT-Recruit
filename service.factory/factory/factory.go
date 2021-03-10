package factory

import (
	"time"

	"github.com/0B1t322/RTUIT-Recruit/pkg/models/product"
	log "github.com/sirupsen/logrus"
)

var productsID chan uint

type ProductCreater interface {
	Create(*product.Product) error
}


type Factory struct {
	productionCapacity 	time.Duration
	productPerCap		uint
	products			[]product.Product
}

type Factorer interface {
	MakeProducts(w Warehouser)
}

func logError(m string, err error) {
	log.WithFields(log.Fields{
		"method": m,
		"err": err,
	}).Error()
}

func NewFactory(
	ProductrionCapacity		time.Duration,
	ProductPerCap			uint,
	prodcuts				[]product.Product,
	p 						ProductCreater,
) *Factory {
	factory := &Factory{
		productionCapacity: ProductrionCapacity,
		productPerCap: ProductPerCap,
		products: prodcuts,
	}
	productsID = make(chan uint, len(prodcuts))
	factory.createProducts(p)
	close(productsID)

	return factory
}

func (f *Factory) createProducts(pc ProductCreater) {
	for i, p := range f.products {
		if err := pc.Create(&p); err != nil {
			logError("createProducts", err)
		} else {
			f.products[i] = p
			productsID <- p.ID
		}
	}
	log.Infof("Products: %v", f.products)
}

func (f *Factory) MakeProducts(w Warehouser) {
	log.Infoln("Start timer")
	timer := time.NewTimer(f.productionCapacity)
	<- timer.C
	f.makeProducts(w)
	log.Infoln("Make products")
}

func (f *Factory) makeProducts(w Warehouser) {
	log.Infof("Products: %v", f.products)
	c := f.productPerCap
	for _, p := range f.products {
		go w.AddProduct(p.ID, c)
	}
}