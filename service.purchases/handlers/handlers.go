package handlers

import (
	"gorm.io/gorm"
	"github.com/0B1t322/RTUIT-Recruit/pkg/models/purchase"
	pc "github.com/0B1t322/RTUIT-Recruit/pkg/controllers/purchase"
	"encoding/json"
	"net/http"
	"time"

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
			"Method": "Get",
			"Error": err,
		},
	).Error()

	w.WriteHeader(http.StatusInternalServerError)
}

// Get return a json purchase according to id in path
//
// if not find purchase fot uid return 404 code and empy body
// 
// 	Answer example:
// {
// 		"ID": 1,
// 		"CreatedAt": "2021-02-18T19:51:42Z",
// 		"UpdatedAt": "2021-02-18T19:51:42Z",
// 		"DeletedAt": null,
// 		"UID": "1",
// 		"buy_date": "2021-02-18T19:51:41.999Z",
// 		"product_name": "prodcut_1",
// 		"cost": 123.2234567
// 	}
func (ph *PurchaseHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uid := vars["uid"]
	
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
// GetAll return a json mass of purchases for uid in path
// 
// if not find purchase for uid return code 404
// Success return code 200
// 
func (ph* PurchaseHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]

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

// Add to db a purchases with uid in path
// 
// if add return code 201 and uint id of added purchase
// 
// body example:
// 	{
// 		"product_name": "product_1",
// 		"cost": 213,
// 	}
func (ph *PurchaseHandler) Add(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]

	d := json.NewDecoder(r.Body)

	p := &purchase.Purchase{}

	if err := d.Decode(p); err != nil {
		logAndWriteAboutInternalError(w, err, "Add")
		return
	}

	p.UID = uid
	p.BuyDate = time.Now()
	if err := ph.c.Create(p); err != nil {
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

// func Update(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	id := vars["id"]

// 	p, err := purchases.Get(id)
// 	if err == purchases.ErrNotFound {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	} else if err != nil {
// 		logAndWriteAboutInternalError(w, err)
// 		return
// 	} 
// }

// Delete a purchase with id in path

// if not found purchase with this id return code 404

// If success retorn code 200
func (ph *PurchaseHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	uid := vars["uid"]

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