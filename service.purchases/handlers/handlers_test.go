package handlers_test

import (
	"crypto/sha512"

	p "github.com/0B1t322/RTUIT-Recruit/pkg/models/purchase"

	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/0B1t322/RTUIT-Recruit/service.purchases/router"
	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
	"github.com/gorilla/mux"
)

func init() {
	m := db.NewManager("root", "root", "127.0.0.1:3306", time.Second)
	db, err := m.OpenDataBase("recruit?parseTime=true")
	if err != nil {
		panic(err)
	}

	r = router.New(db)

	sha := sha512.New()
	sha.Write([]byte("my_secret_key"))
	authHead = "Token " + hex.EncodeToString(sha.Sum(nil))
	println(authHead)
}


var r *mux.Router
var authHead string

func TestFunc_Add(t *testing.T) {
	data, err := json.Marshal(p.Purchase{
		ProductID: 29,
		ShopID: 9,
		Payment: "cash",
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(string(data))

	req := httptest.NewRequest("POST", "/purchases/1", bytes.NewReader(data))
	req.Header.Set("Authorization", authHead)
	w := httptest.NewRecorder()

	r.ServeHTTP(w,req)

	if w.Code != http.StatusCreated {
		t.Log(w.Code)
		t.FailNow()
	}

	var ID uint
	err = json.Unmarshal(w.Body.Bytes(), &ID)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Logf("Added with id: %v\n", ID)

	defer func() {
		req := httptest.NewRequest("DELETE", "/purchases/1/" + fmt.Sprint(ID), nil)
		req.Header.Set("Authorization", authHead)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code !=  http.StatusOK {
			t.Log(w.Code)
			t.FailNow()
		}
	}()
}

func TestFunc_Get(t *testing.T) {
	data, err := json.Marshal(p.Purchase{
		ProductID: 29,
		ShopID: 9,
		Payment: "cash",
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	req := httptest.NewRequest("POST", "/purchases/1", bytes.NewReader(data))
	req.Header.Set("Authorization", authHead)
	w := httptest.NewRecorder()

	r.ServeHTTP(w,req)

	if w.Code != http.StatusCreated {
		t.Log(w.Code)
		t.FailNow()
	}

	var ID uint
	err = json.Unmarshal(w.Body.Bytes(), &ID)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer func() {
		req := httptest.NewRequest("DELETE", "/purchases/1/" + fmt.Sprint(ID), nil)
		req.Header.Set("Authorization", authHead)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code !=  http.StatusOK {
			t.Log(w.Code)
			t.FailNow()
		}
	}()

	t.Logf("Added with id: %v\n", ID)

	req = httptest.NewRequest("GET", fmt.Sprintf("/purchases/1/%v", ID), nil)
	req.Header.Set("Authorization", authHead)
	w = httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Log(w.Code)
		t.FailNow()
	}

	t.Log(w.Body.String())
	// 
}

func TestFunc_Get_NotFound(t *testing.T) {
	req := httptest.NewRequest("GET", fmt.Sprintf("/purchases/10/12"), nil)
	req.Header.Set("Authorization", authHead)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Log(w.Code)
		t.FailNow()
	}

	add := func() (uint, func()) {
		data, err := json.Marshal(p.Purchase{
			ProductID: 29,
			ShopID: 9,
			Payment: "cash",
		})
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
		req := httptest.NewRequest("POST", "/purchases/2", bytes.NewReader(data))
		req.Header.Set("Authorization", authHead)
		w := httptest.NewRecorder()

		r.ServeHTTP(w,req)

		if w.Code != http.StatusCreated {
			t.Log(w.Code)
			t.FailNow()
		}
		var ID uint
		err = json.Unmarshal(w.Body.Bytes(), &ID)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		t.Logf("Added with id: %v\n", ID)

		return ID, (func() {
			req := httptest.NewRequest("DELETE", "/purchases/2/" + fmt.Sprint(ID), nil)
			req.Header.Set("Authorization", authHead)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Log(w.Code)
				t.FailNow()
			}
		})
	}

	id, del := add()
	defer del()

	req = httptest.NewRequest("GET", fmt.Sprintf("/purchases/10/" + fmt.Sprint(id)), nil)
	req.Header.Set("Authorization", authHead)
	w = httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Log(w.Code)
		t.FailNow()
	}
}

func TestFunc_GetAll(t *testing.T) {
	for i := 0; i < 10; i++ {
		data, err := json.Marshal(p.Purchase{
			ProductID: 29,
			ShopID: 9,
			Payment: "cash",
		})
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		req := httptest.NewRequest("POST", "/purchases/1", bytes.NewReader(data))
		req.Header.Set("Authorization", authHead)
		w := httptest.NewRecorder()

		r.ServeHTTP(w,req)

		if w.Code != http.StatusCreated {
			t.Log(w.Code)
			t.FailNow()
		}

		var ID uint
		err = json.Unmarshal(w.Body.Bytes(), &ID)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
		defer func() {
			req := httptest.NewRequest("DELETE", "/purchases/1/" + fmt.Sprint(ID), nil)
			req.Header.Set("Authorization", authHead)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if w.Code !=  http.StatusOK {
				t.Log(w.Code)
				t.FailNow()
			}
		}()
	}

	req := httptest.NewRequest("GET", "/purchases/1", nil)
	req.Header.Set("Authorization", authHead)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Log(w.Code)
		t.FailNow()
	}

	t.Log(w.Body.String())
}

func TestFunc_GetAll_NotFound(t *testing.T) {
	for i := 0; i < 10; i++ {
		data, err := json.Marshal(p.Purchase{
			ProductID: 29,
			ShopID: 9,
			Payment: "cash",
		})

		req := httptest.NewRequest("POST", "/purchases/1", bytes.NewReader(data))
		req.Header.Set("Authorization", authHead)
		w := httptest.NewRecorder()

		r.ServeHTTP(w,req)

		if w.Code != http.StatusCreated {
			t.Log(w.Code)
			t.FailNow()
		}

		var ID uint
		err = json.Unmarshal(w.Body.Bytes(), &ID)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
		defer func() {
			req := httptest.NewRequest("DELETE", "/purchases/1/" + fmt.Sprint(ID), nil)
			req.Header.Set("Authorization", authHead)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if w.Code !=  http.StatusOK {
				t.Log(w.Code)
				t.FailNow()
			}
		}()
	}

	req := httptest.NewRequest("GET", "/purchases/100", nil)
	req.Header.Set("Authorization", authHead)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	t.Log(w.Body.String())
	if w.Code != http.StatusNotFound {
		t.Log(w.Code)
		t.FailNow()
	}
}