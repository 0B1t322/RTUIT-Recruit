package shophandler

import (
	"net/http"
)

type ShopHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Buy(w http.ResponseWriter, r *http.Request)
	GetPurchases(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	AddProduct(w http.ResponseWriter, r *http.Request)
	AddCount(w http.ResponseWriter, r *http.Request)
	CreateShop(w http.ResponseWriter, r *http.Request)
}