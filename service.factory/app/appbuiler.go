package app

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	"github.com/0B1t322/RTUIT-Recruit/pkg/app"
	"github.com/0B1t322/RTUIT-Recruit/pkg/models/product"
	"github.com/0B1t322/RTUIT-Recruit/service.factory/factory"
)

type Config struct {
	ProductuionCapisity	string				`json:"production_capisity"`
	ProductPerCap		uint				`json:"product_per_cap"`
	Products			[]product.Product	`json:"products"`
	ShopNetwork			string				`json:"shop_network"`
	UpdateTime			string				`json:"update_time"`
	DeliveryTime		string				`json:"delivery_time"`
}

func New(cfg *Config, pc factory.ProductCreater) app.App {
	productionCapisity, err := time.ParseDuration(cfg.ProductuionCapisity)
	if err != nil {
		log.Panicf("Failed to build app err on cfg: %v\n", err)
	}
	log.Info("creating factory")
	f := factory.NewFactory(
		productionCapisity,
		cfg.ProductPerCap,
		cfg.Products,
		pc,
	)
	log.Info("factory created")

	updateTime, err := time.ParseDuration(cfg.UpdateTime)
	if err != nil {
		log.Panicf("Failed to build app err on cfg: %v\n", err)
	}

	delivertTime, err := time.ParseDuration(cfg.DeliveryTime)
	if err != nil {
		log.Panicf("Failed to build app err on cfg: %v\n", err)
	}
	log.Info("creating delivery")
	d := factory.NewDelivary(
		cfg.ShopNetwork,
		updateTime,
		delivertTime,
	)
	log.Info("delivery created")

	log.Info("Creating warhouse")
	w := factory.NewWarehouse()
	log.Info("warhouse created")
	
	return &App{
		Factory: f,
		Delivery: d,
		Warehouser: w,
	}
}

func ParseCfg(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	logrus.Infoln(bytes.NewBuffer(data).String())

	cfg := &Config{}

	if err := json.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}