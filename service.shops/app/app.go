package app

import (
	"fmt"
	"net/http"
	"github.com/0B1t322/RTUIT-Recruit/pkg/models/product"
	"github.com/0B1t322/RTUIT-Recruit/pkg/models/shop"
	"github.com/sirupsen/logrus"
	"github.com/0B1t322/RTUIT-Recruit/pkg/middlewares"
	"github.com/0B1t322/RTUIT-Recruit/service.shops/handlers"
	"github.com/0B1t322/RTUIT-Recruit/service.shops/router"
	"github.com/0B1t322/RTUIT-Recruit/pkg/app"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type App struct {
	db *gorm.DB

	port string

	r *mux.Router
}

func New(DB *gorm.DB, Port, PurhacesNetwork string) app.App {
	return &App{
		db: DB,
		port: Port,
		r: router.New(
			handlers.New(DB, PurhacesNetwork),
		),
	}
}

func (a *App) Start() error {
	if err := a.init(); err != nil {
		return err
	}

	return http.ListenAndServe(fmt.Sprintf(":%s", a.port), a.r)
}

func (a *App) init() error {
	middlewares.Logger = logrus.StandardLogger()

	if err := product.AutoMigrate(a.db); err != nil {
		return err
	}

	if err := shop.AutoMigrate(a.db); err != nil {
		return err
	}

	return nil
}