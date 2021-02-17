package purchases

import (
	log "github.com/sirupsen/logrus"
	"github.com/0B1t322/RTUIT-Recruit/pkg/db"
	"time"
	"gorm.io/gorm"
)

func init() {
	if err := db.DB.AutoMigrate(&Purchase{}); err !=  nil {
		log.WithFields(log.Fields{
			"Package": "purchases",
			"Func": "init",
			"Err": err,
		}).Error()
	}
}

type Purchase struct {
	gorm.Model
	BuyDate		time.Time 	`json:"buy_date"`
	ProductName	string		`json:"product_name"`
	Cost		float64		`json:"cost"`
}

func Get(ID string) (*Purchase, error) {
	p := &Purchase{}	
	if err := db.DB.First(p, "ID = ?", ID).Error; err == gorm.ErrRecordNotFound {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return p, nil
}

func Create(p *Purchase) error {
	if err := db.DB.Create(p).Error; err != nil {
		return err
	}

	return nil
}

func (p *Purchase) Update() error {
	if err := db.DB.Updates(p).Error; err != nil {
		return err
	}

	return nil
}

func (p *Purchase) Delete() error {
	if err := db.DB.Delete(p).Error; err != nil {
		return err
	}

	return nil
}