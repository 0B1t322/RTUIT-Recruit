package app

import (
	"github.com/0B1t322/RTUIT-Recruit/pkg/models/purchase"
	"fmt"
	"net/http"

	"github.com/0B1t322/RTUIT-Recruit/service.purchases/router"
	"github.com/gorilla/mux"

	"gorm.io/gorm"
)

type App struct {
	db 		*gorm.DB

	// Maybe set config
	port	string

	r		*mux.Router
}

func New(DB *gorm.DB, Port string) *App {
	return &App{
		db: DB,
		port: Port,
		r: router.New(DB),
	}
}

func (a *App) Start() error {
	if err := a.init(); err  != nil {
		return err
	}
	
	return http.ListenAndServe(fmt.Sprintf(":%s", a.port), a.r)
}

func (a *App) init() error {
	return purchase.AutoMigrate(a.db)
}