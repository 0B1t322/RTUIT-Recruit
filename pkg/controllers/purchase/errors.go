package purchase

import "errors"

var (
	ErrNotFound 		= 	errors.New("Purchase not found")
	ErrInvalidShopID	=  	errors.New("Invalid ShopID: can't find shop")
	ErrInvalidProductID	=	errors.New("Invalid ProductID: can't find product")
	ErrCountNull 		= 	errors.New("Count can't be null")
)