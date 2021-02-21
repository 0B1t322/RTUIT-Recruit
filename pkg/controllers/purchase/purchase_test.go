package purchase_test

import (
	"testing"
	"time"

	"github.com/0B1t322/RTUIT-Recruit/pkg/models/purchase"

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
}

var pc *c.PurchaseController

func TestFunc_Get(t *testing.T) {
	// TODO  create shop
	p := &purchase.Purchase{
		UID: 1,
		ProductName: "Computer",
		Category: "Tech",
		ShopID: 3,
		BuyDate: time.Now(),
		Cost: 89000,
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
}

func TestFunc_Get_NotFound(t *testing.T) {
	_, err := pc.Get(127)
	if err != c.ErrNotFound {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_GetAll(t *testing.T) {
	// TODO  create shop
	for i := 0; i  < 10; i++ {
		p := &purchase.Purchase{
			UID: 1,
			ProductName: "Computer",
			Category: "Tech",
			ShopID: 3,
			BuyDate: time.Now(),
			Cost: 89000 + float64(i),
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

