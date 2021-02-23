package purchase_test

import (
	s "github.com/0B1t322/RTUIT-Recruit/pkg/controllers/shop"
	"encoding/json"
	"testing"
	"time"

	"github.com/0B1t322/RTUIT-Recruit/pkg/models/product"
	"github.com/0B1t322/RTUIT-Recruit/pkg/models/purchase"
	"github.com/0B1t322/RTUIT-Recruit/pkg/models/shop"

	c "github.com/0B1t322/RTUIT-Recruit/pkg/controllers/purchase"
	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
)


func init() {
	manager := db.NewManager("root", "root", "127.0.0.1:3306", time.Second)

	DB, err := manager.OpenDataBase("recruit?parseTime=true")
	if err != nil {
		panic(err)
	}
	
	purchase.AutoMigrate(DB)

	pc = c.New(DB)
	sc = s.New(DB)
}

var pc *c.PurchaseController
var sc *s.ShopController

func TestFunc_Get(t *testing.T) {
	shopModel := &shop.Shop{
		Name: "shop_1",
		Adress: "adress_1",
		PhoneNubmer: "897612334334",
	}

	if err := sc.Create(shopModel); err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer func() {
		if err := sc.Delete(shopModel); err != nil {
			t.Log(err)
			t.FailNow()
		}
	}()

	p := &purchase.Purchase{
		UID: 1,
		Product: product.Product{
			Name: "some_product",
		},
		ShopID: shopModel.ID,
		BuyDate: time.Now(),
	}
	if err := pc.Create(p); err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer pc.Delete(p)


	getP, err := pc.Get(p.ID)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(getP)
	t.Log(getP.Shop)
	t.Log(getP.Product)
	data, err := json.Marshal(getP)
	if err  != nil {
		t.Log(err)
	}

	t.Log(string(data))
}

func TestFunc_Get_NotFound(t *testing.T) {
	_, err := pc.Get(127)
	if err != c.ErrNotFound {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_GetAll(t *testing.T) {
	shopModel := &shop.Shop{
		Name: "shop_2",
		Adress: "adress_1",
		PhoneNubmer: "897612334334",
	}

	if err := sc.Create(shopModel); err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer sc.Delete(shopModel)

	
	for i := 0; i  < 10; i++ {
		p := &purchase.Purchase{
			UID: 1,
			ShopID: shopModel.ID,
			Product: product.Product{
				Name: "Phone",
			},
			BuyDate: time.Now(),
		}

		if err := pc.Create(p); err != nil {
			t.Log(err)
			t.FailNow()
		}
		defer pc.Delete(p)
	}

	ps, err := pc.GetAll(1)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if l := len(ps); l < 10  {
		t.Log(l)
		t.FailNow()
	}

	t.Log(ps)
}

func TestFunc_GetAll_NotFound(t *testing.T) {
	_, err := pc.GetAll(127)
	if err != c.ErrNotFound {
		t.Log(err)
		t.FailNow()
	}
}

