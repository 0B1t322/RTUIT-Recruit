package product

import "gorm.io/gorm"


type Product struct {
	ID				uint	`gorm:"primarykey"`
	Name 			string	`json:"name"`
	Desccription	string	`json:"description"`
	Cost			float64	`json:"cost"`
	Category		string	`json:"category"`
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Product{})
}