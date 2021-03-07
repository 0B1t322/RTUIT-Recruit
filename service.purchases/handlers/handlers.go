package handlers

import (
	sc "github.com/0B1t322/RTUIT-Recruit/pkg/controllers/shop"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

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

// Get return a json purchase according to id in path
//
// if not find purchase fot uid return 404 code and empy body
// 
// 	Answer example:
// {
//     "id": 1,
//     "uid": 1,
//     "shop_id": 9,
//     "shop": {
//         "ID": 9,
//         "name": "cool_shop",
//         "adress": "adress_1",
//         "phone_number": "89991234567"
//     },
//     "buy_date": "2021-02-27T18:32:47.959Z",
//     "product_id": 29,
//     "product": {
//         "ID": 29,
//         "name": "phone_1",
//         "description": "cool phone",
//         "cost": 13000,
//         "category": "Phone"
//     },
//     "payment": "",
//     "count": 1
// }
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
// GetAll return a json mass of purchases for uid in path
// 
// if not find purchase for uid return code 404
// Success return code 200
//
// Answer example:
// [
//     {
//         "id": 1,
//         "uid": 1,
//         "shop_id": 9,
//         "shop": {
//             "ID": 9,
//             "name": "cool_shop",
//             "adress": "adress_1",
//             "phone_number": "89991244567"
//         },
//         "buy_date": "2021-02-27T18:52:20.059Z",
//         "product_id": 29,
//         "product": {
//             "ID": 29,
//             "name": "phone_1",
//             "description": "cool phone",
//             "cost": 13000,
//             "category": "Phone"
//         },
//         "payment": "",
//         "count": 5
//     },
//     {
//         "id": 2,
//         "uid": 1,
//         "shop_id": 9,
//         "shop": {
//             "ID": 9,
//             "name": "cool_shop",
//             "adress": "adress_1",
//             "phone_number": "89991244567"
//         },
//         "buy_date": "2021-02-27T18:52:26.817Z",
//         "product_id": 29,
//         "product": {
//             "ID": 29,
//             "name": "phone_1",
//             "description": "cool phone",
//             "cost": 13000,
//             "category": "Phone"
//         },
//         "payment": "",
//         "count": 1
//     }
// ]
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

// Add to db a purchases with uid in path
// 
// if add return code 201 and uint id of added purchase
// 
// body example:
//	{
//		"product_id": 29,
//		"shop_id": 9,
//		"cost": 999,
//		"count": 1
//	}
func (ph *PurchaseHandler) Add(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	_uid, _ := strconv.ParseUint(vars["uid"], 10, 64)
	uid := uint(_uid)

	d := json.NewDecoder(r.Body)

	p := &purchase.Purchase{}

	if err := d.Decode(p); err != nil {
		logAndWriteAboutInternalError(w, err, "Add")
		return
	}

	p.UID = uid
	p.BuyDate = time.Now()

	if err := ph.c.Create(p); err == pc.ErrInvalidShopID || err == pc.ErrInvalidProductID {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	} else if err == sc.ErrNegCount{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	} else if err == pc.ErrCountNull {
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
// 
// if not found purchase with this id return code 404
// 
// If success retorn code 200
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