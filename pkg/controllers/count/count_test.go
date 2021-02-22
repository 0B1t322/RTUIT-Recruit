package count_test

import (
	"testing"
	"time"

	"github.com/0B1t322/RTUIT-Recruit/pkg/controllers/count"
	m "github.com/0B1t322/RTUIT-Recruit/pkg/models/count"
	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
)

func init() {
	manager := db.NewManager("root", "root", "127.0.0.1:3306", time.Second)

	DB, err := manager.OpenDataBase("recruit?parseTime=true")
	if err != nil {
		panic(err)
	}

	m.AutoMigrate(DB)

	cc = count.New(DB)
}

var cc *count.CountController

func TestFunc_Create(t *testing.T) {
	c := &m.Count{
		ShopID:    14,
		ProductID: 2,
		Count:     40,
	}

	if err := cc.Create(c); err != nil {
		t.Log(err)
		t.FailNow()
	}

	defer func() {
		err := cc.Delete(c)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
	}()
}

func TestFunc_Create_Exist(t *testing.T) {
	c := &m.Count{
		ShopID: 15,
		ProductID: 20,
		Count: 40,
	}

	if err := cc.Create(c); err != nil {
		t.Log(err)
		t.FailNow()
	}

	if err := cc.Create(c); err != count.ErrExist {
		t.Log(err)
		t.FailNow()
	}
	
	defer func() {
		err := cc.Delete(c)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
	}()
}