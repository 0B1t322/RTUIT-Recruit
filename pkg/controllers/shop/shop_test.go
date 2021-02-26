package shop_test

import (
	"encoding/json"
	"testing"
	"time"

	pc "github.com/0B1t322/RTUIT-Recruit/pkg/controllers/product"

	c "github.com/0B1t322/RTUIT-Recruit/pkg/controllers/shop"
	p "github.com/0B1t322/RTUIT-Recruit/pkg/models/product"
	m "github.com/0B1t322/RTUIT-Recruit/pkg/models/shop"

	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
)

func init() {
	manager := db.NewManager("root", "root", "127.0.0.1:3306", time.Second)

	db, err := manager.OpenDataBase("recruit?parseTime=true")
	if err != nil {
		panic(err)
	}

	sc = c.New(db)
	if err := m.AutoMigrate(db); err != nil {
		panic(err)
	}

	if err := m.SetupJoinTable(db); err != nil {
		panic(err)
	}

	pc := pc.New(db)

	pc.Create(&p.Product{
		ID: 29,
		Name: "phone_1",
		Desccription: "cool phone",
		Cost: 13000,
		Category: "Phone",
	})

	pc.Create(&p.Product{
		ID: 30,
		Name: "phone_2",
		Desccription: "coolest phone",
		Cost: 15000,
		Category: "Phone",
	})

}

var sc *c.ShopController

func TestFunc_Create(t *testing.T) {
	s := &m.Shop{
		ShopInfo: m.ShopInfo{
			Name:        "shop_1",
			Adress:      "adress_1",
			PhoneNubmer: "phone_1",
		},
		Products: []p.Product{
			{ID: 29},
		},
	}

	if err := sc.Create(s); err != nil {
		t.Log(err)
		t.FailNow()
	}
	data, _ := json.Marshal(s)
	t.Log(string(data))

	if err := sc.Delete(s); err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_Get(t *testing.T) {
	s := &m.Shop{
		ShopInfo: m.ShopInfo{
			Name:        "shop_1",
			Adress:      "adress_1",
			PhoneNubmer: "phone_1",
		},
		Products: []p.Product{
			{ID: 29}, {ID: 30},
		},
	}

	if err := sc.Create(s); err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer func() {
		if err := sc.Delete(s);err != nil {
			t.Log(err)
			t.FailNow()
		}
	}()

	if data, _ := json.Marshal(s); true {
		t.Log(string(data))
	}
	t.Log("Get")
	getS, err := sc.Get(s.ID)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	data, _ := json.Marshal(getS)
	t.Log(string(data))
}

func TestFunc_Get_NotFound(t *testing.T) {
	_, err := sc.Get(1)
	if err != c.ErrNotFound {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_Update(t *testing.T) {
	s := &m.Shop{
		ShopInfo: m.ShopInfo{
			Name:        "shop_1",
			Adress:      "adress_1",
			PhoneNubmer: "phone_1",
		},
		Products: []p.Product{
			{ID: 29}, {ID: 30},
		},
	}

	if err := sc.Create(s); err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer func() {
		if err := sc.Delete(s);err != nil {
			t.Log(err)
			t.FailNow()
		}
	}()

	if err := sc.Update(s); err !=  nil {
		t.Log(err)
		t.FailNow()
	}
	data, _ := json.Marshal(s)
	t.Log(string(data))
}

func TestFunc_AddCount(t *testing.T) {
	s := &m.Shop{
		ShopInfo: m.ShopInfo{
			Name:        "shop_1",
			Adress:      "adress_1",
			PhoneNubmer: "phone_1",
		},
		Products: []p.Product{
			{ID: 29}, {ID: 30},
		},
	}
	if err := sc.Create(s); err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer func() {
		if err := sc.Delete(s);err != nil {
			t.Log(err)
			t.FailNow()
		}
	}()
	
	if err := sc.AddCount(s.ID, 29, 10); err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_AddCount_ProductNotFound(t *testing.T) {
	s := &m.Shop{
		ShopInfo: m.ShopInfo{
			Name:        "shop_1",
			Adress:      "adress_1",
			PhoneNubmer: "phone_1",
		},
		Products: []p.Product{
			{ID: 29}, {ID: 30},
		},
	}
	if err := sc.Create(s); err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer func() {
		if err := sc.Delete(s);err != nil {
			t.Log(err)
			t.FailNow()
		}
	}()
	
	if err := sc.AddCount(s.ID, 1, 10); err != c.ErrProductNotFound {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_SubCount(t *testing.T) {
	s := &m.Shop{
		ShopInfo: m.ShopInfo{
			Name:        "shop_1",
			Adress:      "adress_1",
			PhoneNubmer: "phone_1",
		},
		Products: []p.Product{
			{ID: 29}, {ID: 30},
		},
	}
	if err := sc.Create(s); err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer func() {
		if err := sc.Delete(s);err != nil {
			t.Log(err)
			t.FailNow()
		}
	}()

	if err := sc.AddCount(s.ID, 29,1); err != nil {
		t.Log(err)
		t.FailNow()
	}

	if err := sc.SubCount(s.ID, 29,1); err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_SubCount_ErrNegCount(t *testing.T) {
	s := &m.Shop{
		ShopInfo: m.ShopInfo{
			Name:        "shop_1",
			Adress:      "adress_1",
			PhoneNubmer: "phone_1",
		},
		Products: []p.Product{
			{ID: 29}, {ID: 30},
		},
	}
	if err := sc.Create(s); err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer func() {
		if err := sc.Delete(s);err != nil {
			t.Log(err)
			t.FailNow()
		}
	}()

	if err := sc.SubCount(s.ID, 29,1); err != c.ErrNegCount {
		t.Log(err)
		t.FailNow()
	}
}