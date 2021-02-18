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

type PurchaseHandler struct {
	c *pc.PurchaseController
}

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

func (ph *PurchaseHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	p, err := ph.c.Get(id)
	if err == pc.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		logAndWriteAboutInternalError(w, err, "Get")
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

func (ph *PurchaseHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	p, err := ph.c.Get(id)
	if err == pc.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		logAndWriteAboutInternalError(w, err, "Delete")
		return
	}

	if err := ph.c.Delete(p); err != nil {
		logAndWriteAboutInternalError(w, err, "Delete")
		return
	}

	w.WriteHeader(http.StatusOK)
}