package main

import (
	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
	"flag"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"

	pc "github.com/0B1t322/RTUIT-Recruit/pkg/controllers/product"
	"github.com/0B1t322/RTUIT-Recruit/pkg/models/product"

	"github.com/0B1t322/RTUIT-Recruit/service.factory/app"
	"github.com/0B1t322/RTUIT-Recruit/service.factory/factory"
)

const ShopNetwork = "http://localhost:8082"

var creater *pc.ProductController

var done chan os.Signal

func init() {
	m := db.NewManager("root", "root", "127.0.0.1:3306", 15*time.Second)
	db, err := m.OpenDataBase("recruit?parseTime=true")
	if err != nil {
		panic(err)
	}
	creater = pc.New(db)

	done = make(chan os.Signal)
}

func main() {
	flag.Parse()
	logrus.SetLevel(logrus.ErrorLevel)
	app := app.New(
		factory.NewFactory(
			15 * time.Second,
			3,
			[]product.Product{
				{Name: "iPhone X 64 GB", Desccription: "New iphone with 64 gb", Cost: 57000.0, Category: "Phones"},
				{Name: "iPhone X 128 GB", Desccription: "New iphone with 128 gb", Cost: 77000.0, Category: "Phones"},
				{Name: "iPhone X 256 GB", Desccription: "New iphone with 256 gb", Cost: 97000.0, Category: "Phones"},
			},
			creater,
		),
		factory.NewDelivary(
			ShopNetwork,
			10 * time.Second,
			5 * time.Second,
		),
		factory.NewWarehouse(),
	)

	logrus.Infoln("Start service factory")
	if err := app.Start(); err != nil {
		logrus.Panic(err)
	}

	signal.Notify(done, os.Interrupt, syscall.SIGTERM)
	select {
	case <- done:
		logrus.Infoln("Shuting down")
		os.Exit(0)
	}
}