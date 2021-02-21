package shop

import "gorm.io/gorm"

type Shop struct {
	ID			uint	`gorm:"primarykey"`
	Name		string 	`json:"name"`
	Adress		string	`json:"adress"`
	PhoneNubmer	string	`json:"phone_number"`
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Shop{})
}