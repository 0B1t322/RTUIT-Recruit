package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	pm "github.com/0B1t322/RTUIT-Recruit/pkg/models/purchase"

	log "github.com/sirupsen/logrus"

	sc "github.com/0B1t322/RTUIT-Recruit/pkg/controllers/shop"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type ShopHandler struct {
	c 				*sc.ShopController
	purhacesNetwork	string
	AuthKey			string
}

func New(db *gorm.DB) *ShopHandler {
	return &ShopHandler{
		c: sc.New(db),
	}
}

func logAndWriteAboutInternalError(w http.ResponseWriter, err error, m string) {
	log.WithFields(
		log.Fields{
			"Package": "handlers",
			"Method": m,
			"Error": err,
		},
	).Error()

	w.WriteHeader(http.StatusInternalServerError)
}

func (sp *ShopHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := getShopID(vars)

	shop, err := sp.c.Get(id)
	if err == sc.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	} else if err != nil {
		logAndWriteAboutInternalError(w, err, "Get")
		return
	}

	data, err := json.Marshal(shop)
	if err != nil {
		logAndWriteAboutInternalError(w, err, "Get")
		return
	}

	w.Write(data)
	w.WriteHeader(http.StatusOK)
}

func (sp *ShopHandler) Buy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := getShopID(vars)
	productID := getProductID(vars)

	p := &pm.Purchase{}

	d := json.NewDecoder(r.Body)
	if err := d.Decode(p); err != nil {
		logAndWriteAboutInternalError(w, err, "Buy")
		return
	}
	
	p.ShopID = id
	p.ProductID = productID

	if err := sp.buy(p); err !=  nil {
		logAndWriteAboutInternalError(w, err, "Buy")
	}

	w.WriteHeader(http.StatusOK)
}

func getShopID(v map[string]string) uint {
	return getUINT(v, "id")
}

func getProductID(v map[string]string) uint {
	return getUINT(v, "pid")
}

func getUINT(v map[string]string, name string) uint {
	get := v[name]

	_uint, _ := strconv.ParseUint(get, 10, 64)

	return uint(_uint)
}

func (sp *ShopHandler) buy(p *pm.Purchase) error {
	url := fmt.Sprintf("%s/purchases/%v", sp.purhacesNetwork, p.UID)
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}

	body := bytes.NewReader(data)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", sp.AuthKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("Unexpected code: %v", resp.StatusCode)
	}

	return nil
}