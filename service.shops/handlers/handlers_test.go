package handlers_test

import (
	"bytes"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	h "github.com/0B1t322/RTUIT-Recruit/service.shops/handlers"

	"github.com/0B1t322/RTUIT-Recruit/pkg/controllers/purchase"
	"github.com/0B1t322/RTUIT-Recruit/pkg/middlewares"

	pm "github.com/0B1t322/RTUIT-Recruit/pkg/models/purchase"
	"github.com/0B1t322/RTUIT-Recruit/pkg/models/shop"

	"github.com/0B1t322/RTUIT-Recruit/service.shops/router"

	"github.com/0B1t322/distanceLearningWebSite/pkg/db"

	"github.com/gorilla/mux"
)

func init() {
	manager := db.NewManager("root", "root", "127.0.0.1:3306", time.Second)
	db, err := manager.OpenDataBase("recruit?parseTime=true")
	if err != nil {
		panic(err)
	}
	const purNet = "http://localhost:8081"

	route = router.New(h.New(db, purNet))
}

var route *mux.Router

func TestFunc_Buy(t *testing.T) {
	p := &pm.Purchase{
		UID:     1,
		Payment: "cash",
		Count:   100,
	}

	data, err := json.Marshal(p)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	body := bytes.NewReader(data)

	req, err := http.NewRequest("PUT", "/shops/1/2", body)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	w := httptest.NewRecorder()

	route.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Log(w.Code)
		t.Log(w.Body.String())
		t.FailNow()
	}

	

}

func TestFunc_Buy_CountNull(t *testing.T) {
	p := &pm.Purchase{
		UID:     1,
		Payment: "cash",
		Count:   0,
	}

	data, err := json.Marshal(p)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	//
	//
	//

	body := bytes.NewReader(data)

	req, err := http.NewRequest("PUT", "/shops/1/2", body)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	w := httptest.NewRecorder()

	route.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest{
		t.Log(w.Code)
		t.Log(w.Body.String())
		t.FailNow()
	}
}

func TestFunc_Buy_CountNeg(t *testing.T) {
	p := &pm.Purchase{
		UID:     1,
		Payment: "cash",
		Count:   101,
	}

	data, err := json.Marshal(p)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	//
	//
	//

	body := bytes.NewReader(data)

	req, err := http.NewRequest("PUT", "/shops/1/2", body)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	w := httptest.NewRecorder()

	route.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest{
		t.Log(w.Code)
		t.Log(w.Body.String())
		t.FailNow()
	}

	t.Log(w.Body.String())
}


func TestFunc_Buy_NotFound_InvalidShopID(t *testing.T) {
	p := &pm.Purchase{
		UID:     1,
		Payment: "cash",
		Count:   12,
	}

	data, err := json.Marshal(p)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	//
	//
	//

	body := bytes.NewReader(data)

	req, err := http.NewRequest("PUT", "/shops/120/2", body)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	w := httptest.NewRecorder()

	route.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Log(w.Code)
		t.Log(w.Body.String())
		t.FailNow()
	}

	if w.Body.String() != purchase.ErrInvalidShopID.Error() {
		t.Log(w.Body.String())
		t.FailNow()
	}

}

func TestFunc_Buy_NotFound_InvalidProductID(t *testing.T) {
	p := &pm.Purchase{
		UID:     1,
		Payment: "cash",
		Count:   12,
	}

	data, err := json.Marshal(p)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	//
	//
	//

	body := bytes.NewReader(data)

	req, err := http.NewRequest("PUT", "/shops/1/120", body)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	w := httptest.NewRecorder()

	route.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Log(w.Code)
		t.Log(w.Body.String())
		t.FailNow()
	}

	if w.Body.String() != purchase.ErrInvalidProductID.Error() {
		t.Log(w.Body.String())
		t.FailNow()
	}

}

func TestFunc_GetPurchases(t *testing.T) {
	req, err := http.NewRequest("GET", "/shops/purchases/1", nil)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	w := httptest.NewRecorder()

	route.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Log(w.Code)
		t.Log(w.Body.String())
		t.FailNow()
	}

	t.Log(w.Body.String())
}

func TestFunc_GetPurchases_NotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/shops/purchases/2", nil)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	w := httptest.NewRecorder()

	route.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Log(w.Code)
		t.Log(w.Body.String())
		t.FailNow()
	}
}

func TestFunc_Get(t *testing.T) {
	r, err := http.NewRequest("GET", "/shops/1", nil)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	w := httptest.NewRecorder()

	route.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Log(w.Code)
		t.FailNow()
	}

	t.Log(w.Body.String())
}

func TestFunc_Get_NotFound(t *testing.T) {
	r, err := http.NewRequest("GET", "/shops/127", nil)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	w := httptest.NewRecorder()

	route.ServeHTTP(w, r)

	if w.Code != http.StatusNotFound{
		t.Log(w.Code)
		t.FailNow()
	}
}

func TestFunc_GetAll(t *testing.T) {
	r, err := http.NewRequest("GET", "/shops", nil)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	w := httptest.NewRecorder()

	route.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Log(w.Code)
		t.FailNow()
	}

	t.Log(w.Body.String())
}

func TestFunc_AddProduct(t *testing.T) {
	// REQUIRE product with id 3 in db
	r, err := http.NewRequest("POST", "/shops/1/3", nil)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	setAuthHeader(r)
	w := httptest.NewRecorder()

	route.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Log(w.Code)
		t.Log(w.Body.String())
		t.FailNow()
	}
}

func TestFunc_AddProduct_NotFound(t *testing.T) {
	r, err := http.NewRequest("POST", "/shops/2/3", nil)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	setAuthHeader(r)
	w := httptest.NewRecorder()

	route.ServeHTTP(w, r)

	if w.Code != http.StatusNotFound {
		t.Log(w.Code)
		t.Log(w.Body.String())
		t.FailNow()
	}
}

func TestFunc_AddProduct_BadRequest(t *testing.T) {
	r, err := http.NewRequest("POST", "/shops/1/4", nil)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	setAuthHeader(r)
	w := httptest.NewRecorder()

	route.ServeHTTP(w, r)

	if w.Code != http.StatusBadRequest {
		t.Log(w.Code)
		t.Log(w.Body.String())
		t.FailNow()
	}
}


func TestFunc_AddCount(t *testing.T) {
	r, err := http.NewRequest("PUT", "/shops/1/3/2", nil)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	setAuthHeader(r)
	w := httptest.NewRecorder()

	route.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Log(w.Code)
		t.Log(w.Body.String())
		t.FailNow()
	}
}

func setAuthHeader(req *http.Request) {
	sha := sha512.New()
	sha.Write([]byte(middlewares.SecretKey))

	data := sha.Sum(nil)

	token := fmt.Sprintf("Token %s", hex.EncodeToString(data))

	req.Header.Add("Authorization", token)
}

func TestFunc_CreateShop(t *testing.T) {
	s := &shop.ShopInfo{
		Adress: "some_adress",
		Name: "some_shop",
		PhoneNubmer: "8991234567",
	}
	data, err := json.Marshal(s)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	b := bytes.NewReader(data)

	r, err := http.NewRequest("POST", "/shops/", b)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	setAuthHeader(r)

	w := httptest.NewRecorder()

	route.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Log(w.Code)
		t.FailNow()
	}

	t.Log(w.Body.String())
}

func TestFunc_CreateShop_BadReq(t *testing.T) {
	s := &shop.ShopInfo{
		Adress: "some_adress",
		Name: "some_shop",
		PhoneNubmer: "8991234567",
	}
	data, err := json.Marshal(s)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	b := bytes.NewReader(data)

	r, err := http.NewRequest("POST", "/shops/", b)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	setAuthHeader(r)

	w := httptest.NewRecorder()

	route.ServeHTTP(w, r)

	if w.Code != http.StatusBadRequest {
		t.Log(w.Code)
		t.FailNow()
	}
}