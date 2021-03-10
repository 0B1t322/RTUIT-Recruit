package factory

import (
	"bytes"
	"crypto/sha512"
	"encoding/hex"
	"io/ioutil"

	"github.com/0B1t322/RTUIT-Recruit/pkg/middlewares"
	log "github.com/sirupsen/logrus"

	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/jinzhu/copier"

	m "github.com/0B1t322/RTUIT-Recruit/pkg/models/shop"

	"github.com/0B1t322/RTUIT-Recruit/service.factory/hash"
)

type Delivery struct {
	ShopNetwork			string
	dests 				map[uint]hash.HashUint
	UpdateTime			time.Duration
	mu					sync.RWMutex
	DeliverTime			time.Duration
}

type Deliverer interface {
	UpdateShops()
	Deliver(w Warehouser)
}

func NewDelivary(
	ShopNetwork string,
	UpdateTime 	time.Duration,
	DeliverTime	time.Duration,
) *Delivery {
	return &Delivery{
		ShopNetwork: ShopNetwork,
		UpdateTime: UpdateTime,
		dests: make(map[uint]hash.HashUint),
		DeliverTime: DeliverTime,
	}
}

func(d *Delivery) UpdateShops() {
	log.Infoln("Start timer for Update shops")
	timer := time.NewTimer(d.UpdateTime)
	<- timer.C
	d.updateShops()
	log.Infoln("Updated Shop")
	log.Infof("map : %v", d.dests)
}

func (d *Delivery) updateShops() {
	url := fmt.Sprintf("%s/shops", d.ShopNetwork)

	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logError("updateShops", err)
		return
	}

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		logError("updateShops", err)
		return
	}
	defer resp.Body.Close()

	var shops []m.Shop
	{
		d := json.NewDecoder(resp.Body)
		if err := d.Decode(&shops); err != nil {
			logError("updateShops", err)
			return
		}
	}

	d.mu.RLock()
	for _, shop := range shops {
		for _, shopProduct := range shop.ShopProducts {
			h, find := d.dests[shopProduct.ProductID]
			if !find {
				d.dests[shopProduct.ProductID] = make(hash.HashUint)
				h = d.dests[shopProduct.ProductID]
			}

			h.Add(shopProduct.ShopID)
		}
	}
	d.mu.RUnlock()
}

func (d *Delivery) Deliver(w Warehouser) {
	log.Infoln("Start timer for deliver")
	timer := time.NewTimer(d.DeliverTime)
	<- timer.C
	log.Infoln("Start deliver")
	var _dest map[uint]hash.HashUint
	d.mu.Lock()
	log.Infoln("Start copy")
	if err := copier.Copy(&_dest, &d.dests); err != nil {
		log.Errorf("Error on copy %v\n", err)
	}
	log.Infoln("End copy")
	d.mu.Unlock()

	for PID, h := range _dest {
		var shopsID []uint
		shopsID = h.GetKeys()
		for _, shopID := range shopsID {
			log.Infof("start gorutine PID:%v SID:%v", PID, shopID)
			go d.deliver(PID, shopID, w)
		}

	}
	log.Infoln("End deliver")
}

func (d *Delivery) deliver(PID, SID uint, w Warehouser) {
	log.Infoln("Start select")
	select {
	case takeCount := <- w.TakeProduct(PID):
		log.Infof("Take count: %v\n", takeCount)
		d.deliveryToShop(PID, SID, takeCount, w)
	case <- time.After(time.Second):
		log.Infoln("Timeout")
		return
	}
}

func (d *Delivery) deliveryToShop(PID uint, SID uint, count uint8, w Warehouser) {
	url := fmt.Sprintf("%s/shops/%v/%v/%v", d.ShopNetwork, SID, PID, count)
	log.Infoln(url)

	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		logError("deliveryToShop", err)
		d.sendBack(PID, w)
		return
	}

	setAuthHeader(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logError("deliveryToShop", err)
		d.sendBack(PID, w)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		func () {
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logError("deliveryToShop", err)
				data = []byte("Empty")
			}
			logError("deliveryToShop", fmt.Errorf("Status Code: %v, Body: %s", resp.StatusCode, bytes.NewBuffer(data).String()))
		}()
		d.sendBack(PID, w)
		return
	}

	log.Infof("Deliver to PID: %v to ShopID: %v\n", PID, SID)
}

func (d *Delivery) sendBack(PID uint ,w Warehouser) {
	w.AddProduct(PID, 1)
}

func setAuthHeader(req *http.Request) {
	sha := sha512.New()
	sha.Write([]byte(middlewares.SecretKey))

	data := sha.Sum(nil)

	token := fmt.Sprintf("Token %s", hex.EncodeToString(data))

	req.Header.Add("Authorization", token)
}