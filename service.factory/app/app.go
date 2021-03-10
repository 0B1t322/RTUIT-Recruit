package app

import (
	"github.com/0B1t322/RTUIT-Recruit/pkg/app"
	"github.com/0B1t322/RTUIT-Recruit/service.factory/factory"
)

type App struct {
	factory		factory.Factorer
	delivery	factory.Deliverer
	warehouser	factory.Warehouser
}

func New(
	factory		factory.Factorer,
	delivery	factory.Deliverer,
	warehouser	factory.Warehouser,
) app.App {
	return &App{
		factory: factory,
		delivery: delivery,
		warehouser: warehouser,
	}
}

func (a *App) Start() error {
	go func() {
		for {
			a.factory.MakeProducts(a.warehouser)
		}
	}()

	go func() {
		for {
			a.delivery.UpdateShops()
		}
	}()

	go func() {
		for {
			a.delivery.Deliver(a.warehouser)
		}
	}()

	return nil
}