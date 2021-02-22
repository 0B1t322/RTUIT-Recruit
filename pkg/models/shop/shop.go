package shop

import (
	"github.com/0B1t322/RTUIT-Recruit/pkg/models/count"
	"github.com/0B1t322/RTUIT-Recruit/pkg/models/product"
	"gorm.io/gorm"
)

// TODO gorm tags for product and counts

type Shop struct {
	ID			uint				`gorm:"primarykey"`
	Name		string 				`json:"name"`
	Adress		string				`json:"adress"`
	PhoneNubmer	string				`json:"phone_number"`

	// here all existing  products
	Products	[]product.Product	`json:"products" gorm:"-"`
	// here product count
	Counts		[]count.Count		`json:"counts" gorm:"-"`
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Shop{})
}