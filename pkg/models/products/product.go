package products

type Product struct {
	ID				uint	`gorm:"primarykey"`
	Name 			string	`json:"name"`
	Desccription	string	`json:"description"`
	Cost			float64	`json:"cost"`
}