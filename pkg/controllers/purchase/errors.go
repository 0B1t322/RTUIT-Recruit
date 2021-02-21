package purchase

import "errors"

var (
	ErrNotFound 		= 	errors.New("Not found")
	ErrInvalidShopID	=  	errors.New("Invalid ShopID: can't find shop")
)