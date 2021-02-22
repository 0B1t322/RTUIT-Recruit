package product_test

import (
	"encoding/json"
	"testing"
	"time"

	m "github.com/0B1t322/RTUIT-Recruit/pkg/models/product"

	c "github.com/0B1t322/RTUIT-Recruit/pkg/controllers/product"

	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
)

func init() {
	m := db.NewManager("root","root", "127.0.0.1:3306", time.Second)

	db, err := m.OpenDataBase("recruit?parseTime=true")
	if err != nil {
		panic(err)
	}

	pc  = c.New(db)
}

var pc *c.ProductController

func TestFunc_Create(t *testing.T) {
	p := &m.Product{
		Name: "product_1",
		Desccription: "desc_1",
		Cost: 999,
		Category: "category_1",
	}
	
	if err := pc.Create(p); err != nil{
		t.Log(err)
		t.FailNow()
	}

	defer pc.Delete(p)
}

func TestFunc_Get(t *testing.T) {
	p := &m.Product{
		Name: "product_1",
		Desccription: "desc_1",
		Cost: 999,
		Category: "category_1",
	}
	if err := pc.Create(p); err != nil{
		t.Log(err)
		t.FailNow()
	}
	defer pc.Delete(p)

	if getP, err := pc.Get(p.ID); err != nil {
		t.Log(err)
		t.FailNow()
	} else {
		data, err := json.Marshal(getP)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
		t.Log(string(data))
	}
}

func TestFunc_Get_NotFound(t *testing.T) {
	if _, err := pc.Get(12); err != c.ErrNotFound {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_Update(t *testing.T) {
	p := &m.Product{
		Name: "product_1",
		Desccription: "desc_1",
		Cost: 999,
		Category: "category_1",
	}
	if err := pc.Create(p); err != nil{
		t.Log(err)
		t.FailNow()
	}
	defer pc.Delete(p)

	p.Cost = 1000
	if err := pc.Update(p); err != nil {
		t.Log(err)
		t.FailNow()
	}
}