package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	sc "github.com/0B1t322/RTUIT-Recruit/pkg/controllers/shop"

	pc "github.com/0B1t322/RTUIT-Recruit/pkg/controllers/purchase"
	"github.com/0B1t322/RTUIT-Recruit/pkg/models/purchase"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

// PurchaseHandler present a struct with controller
type PurchaseHandler struct {
	c *pc.PurchaseController
}

// New .....
func New(db *gorm.DB) *PurchaseHandler {
	return &PurchaseHandler{
		c: pc.New(db),
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

// Get
// @Summary Get purchases
// @Description get purchase by id
// @ID get-purchase-by-id
// @Produce  json
// @Param   uid      path   int     true  "ID of the user"
// @Param   id      path   int     true  "ID of the purchase"
// @Success 200 {object} purchase.Purchase
// @Failure 404 {string} string "Not found"
// @Router /purchases/{uid}/{id} [get]
func (ph *PurchaseHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_id, _ := strconv.ParseUint(vars["id"], 10, 64)
	_uid, _ := strconv.ParseUint(vars["uid"], 10, 64)
	
	id := uint(_id)
	uid := uint(_uid)

	p, err := ph.c.Get(id)
	if err == pc.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		logAndWriteAboutInternalError(w, err, "Get")
		return
	}

	if p.UID != uid {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := json.Marshal(p)
	if err != nil {
		logAndWriteAboutInternalError(w, err, "Get")
		return
	}

	w.Write(data)
	w.WriteHeader(http.StatusOK)
}

// GetAll
// @Summary Get All purchases
// @Description get All purchase by uid
// @ID get-all-purchase-by-uid
// @Produce  json
// @Param   uid      path   int     true  "ID of the user"
// @Success 200 {array} purchase.Purchase
// @Failure 404 {string} string "Not found"
// @Router /purchases/{uid} [get]
func (ph* PurchaseHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	_uid, _ := strconv.ParseUint(vars["uid"], 10, 64)
	uid := uint(_uid)

	ps, err := ph.c.GetAll(uid)
	if err == pc.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		logAndWriteAboutInternalError(w, err, "GetAll")
		return
	}

	data, err := json.Marshal(ps)
	if err !=  nil {
		logAndWriteAboutInternalError(w,err,"GetAll")
		return
	}

	w.Write(data)
	w.WriteHeader(http.StatusOK)
}

type addBody struct {
	ProductID uint 	`json:"product_id"`
	ShopID uint 	`json:"shop_id"`
	Count uint 		`json:"count"`
	Payment string 	`json:"payment"`
}

// Add
// @Summary Add purchase
// @Description Add purchase
// @ID add-purchase-to-uid
// @Accept json
// @Produce  json
// @Param   uid      path   int     true  "ID of the user"
// @Param   purchase body addBody true "purchase info"
// @Success 201 {integer} integer "id of purchase"
// @Failure 404 {string} string "Not found shop"
// @Failure 404 {string} string "Not found product"
// @Failure 400 {string} string "body is null"
// @Failure 400 {string} string "NegCount"
// @Failure 400 {string} string "count can't be null"
// @Router /purchases/{uid} [post]
func (ph *PurchaseHandler) Add(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	_uid, _ := strconv.ParseUint(vars["uid"], 10, 64)
	uid := uint(_uid)

	d := json.NewDecoder(r.Body)

	p := &purchase.Purchase{}

	if err := d.Decode(p); err == io.EOF {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Body is null")
		return
	} else if err != nil {
		logAndWriteAboutInternalError(w, err, "Add")
		return
	}

	p.UID = uid
	p.BuyDate = time.Now()

	if err := ph.c.Create(p); err == pc.ErrInvalidShopID || err == pc.ErrInvalidProductID {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	} else if err == sc.ErrNegCount || err == pc.ErrCountNull{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	} else if err != nil {
		logAndWriteAboutInternalError(w, err, "Add")
		return
	}

	w.WriteHeader(http.StatusCreated)
	ID := p.ID

	data, err := json.Marshal(ID)
	if err != nil {
		logAndWriteAboutInternalError(w, err, "Add")
		return
	}

	w.Write(data)
}

// Delete
// @Summary Delete purchase
// @Description delete purchase for user by id
// @ID delete-purchase-by-id
// @Param   uid      path   int     true  "ID of the user"
// @Param   id      path   int     true  "ID of the purchase"
// @Success 200
// @Failure 404 {string} string "Not found"
// @Router /purchases/{uid}/{id} [delete]
func (ph *PurchaseHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	_id, _ := strconv.ParseUint(vars["id"], 10, 64)
	_uid, _ := strconv.ParseUint(vars["uid"], 10, 64)
	
	id := uint(_id)
	uid := uint(_uid)

	p, err := ph.c.Get(id)
	if err == pc.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		logAndWriteAboutInternalError(w, err, "Delete")
		return
	}

	if p.UID != uid {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := ph.c.Delete(p); err != nil {
		logAndWriteAboutInternalError(w, err, "Delete")
		return
	}

	w.WriteHeader(http.StatusOK)
}