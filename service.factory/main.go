package main

import (
	"github.com/0B1t322/RTUIT-Recruit/pkg/db"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	pc "github.com/0B1t322/RTUIT-Recruit/pkg/controllers/product"

	"github.com/0B1t322/RTUIT-Recruit/service.factory/app"
)

var creater *pc.ProductController

var done chan os.Signal

func init() {
	// m := db.NewManager("root", "root", "127.0.0.1:3306", 15*time.Second)
	db, err := db.DBManager.OpenDataBase("recruit?parseTime=true")
	if err != nil {
		panic(err)
	}
	creater = pc.New(db)

	done = make(chan os.Signal)
}

func main() {
	flag.Parse()
	// logrus.SetLevel(logrus.ErrorLevel)
	cfg, err := app.ParseCfg("config.json")
	if err != nil {
		logrus.Panicf("Failed to parse config: %v", err)
	}

	app := app.New(
		cfg,
		creater,
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