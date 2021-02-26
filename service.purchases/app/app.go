package app

import (
	"github.com/sirupsen/logrus"
	"github.com/0B1t322/RTUIT-Recruit/pkg/middlewares"
	"github.com/0B1t322/RTUIT-Recruit/pkg/models/purchase"
	"fmt"
	"net/http"

	"github.com/0B1t322/RTUIT-Recruit/service.purchases/router"
	"github.com/gorilla/mux"

	"gorm.io/gorm"
)

// App present a struct with db port and router
// API of app:
// 	GET /purchases/:uid/:id return a purchase
// 	GET /purchases/:uid return all purchase for UID
// 	POST /purchases/:uid add purchase for current user
// 	DELETE /purchases/:uid/:id delete a purchase with this id
type App struct {
	db 		*gorm.DB

	// Maybe set config
	port	string

	r		*mux.Router
}

// New return a pointer for new app
func New(DB *gorm.DB, Port string) *App {
	return &App{
		db: DB,
		port: Port,
		r: router.New(DB),
	}
}

// Start app
func (a *App) Start() error {
	if err := a.init(); err  != nil {
		return err
	}

	return http.ListenAndServe(fmt.Sprintf(":%s", a.port), a.r)
}

func (a *App) init() error {
	middlewares.Logger = logrus.StandardLogger()
	return purchase.AutoMigrate(a.db)
}