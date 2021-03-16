package app

import (
	"github.com/0B1t322/RTUIT-Recruit/service.factory/factory"
)

type App struct {
	Factory    factory.Factorer
	Delivery   factory.Deliverer
	Warehouser factory.Warehouser
}

func (a *App) Start() error {
	go func() {
		for {
			a.Factory.MakeProducts(a.Warehouser)
		}
	}()

	go func() {
		for {
			a.Delivery.UpdateShops()
		}
	}()

	go func() {
		for {
			a.Delivery.Deliver(a.Warehouser)
		}
	}()

	return nil
}
