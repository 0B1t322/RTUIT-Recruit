package shop_test

import (
	"encoding/json"
	"testing"
	"time"

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

}

var sc *c.ShopController

func TestFunc_Create(t *testing.T) {
	s := &m.Shop{
		Name:        "shop_1",
		Adress:      "adress_1",
		PhoneNubmer: "phone_1",
		Products: []p.Product{
			{Name: "name_3"},
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
