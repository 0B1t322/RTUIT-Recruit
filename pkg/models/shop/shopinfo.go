package shop

type ShopInfo struct {
	ID			uint				`gorm:"primarykey"`
	Name		string 				`json:"name" gorm:"unique"`
	Adress		string				`json:"adress"`
	PhoneNubmer	string				`json:"phone_number"`
}