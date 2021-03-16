package handlers

import (
	"io"
	"bytes"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	u "net/url"
	"strconv"

	"github.com/0B1t322/RTUIT-Recruit/pkg/controllers/product"
	"github.com/0B1t322/RTUIT-Recruit/pkg/controllers/shop"
	m "github.com/0B1t322/RTUIT-Recruit/pkg/models/shop"

	pc "github.com/0B1t322/RTUIT-Recruit/pkg/controllers/purchase"

	"github.com/0B1t322/RTUIT-Recruit/pkg/middlewares"
	pm "github.com/0B1t322/RTUIT-Recruit/pkg/models/purchase"

	log "github.com/sirupsen/logrus"

	sc "github.com/0B1t322/RTUIT-Recruit/pkg/controllers/shop"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type ShopHandler struct {
	c 				*sc.ShopController
	PurhacesNetwork	string
}

func New(db *gorm.DB, PurhacesNetwork string) *ShopHandler {
	return &ShopHandler{
		c: sc.New(db),
		PurhacesNetwork: PurhacesNetwork,
	}
}

func logAndWriteAboutInternalError(w http.ResponseWriter, err error, m string) {
	log.WithFields(
		log.Fields{
			"Package": "handlers",
			"Method": m,
			"Error": err,
		},
	).Error()

	w.WriteHeader(http.StatusInternalServerError)
}
// Get
// @Summary Get a shop
// @Description get shop by ID
// @ID get-shop-by-int
// @Produce  json
// @Param   id      path   int     true  "ID of the shop"
// @Success 200 {object} m.Shop
// @Failure 404 {string} string "Shop don't find"
// @Router /shops/{id} [get]
func (sp *ShopHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := getShopID(vars)

	shop, err := sp.c.Get(id)
	if err == sc.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	} else if err != nil {
		logAndWriteAboutInternalError(w, err, "Get")
		return
	}

	data, err := json.Marshal(shop)
	if err != nil {
		logAndWriteAboutInternalError(w, err, "Get")
		return
	}

	w.Write(data)
	w.WriteHeader(http.StatusOK)
}

type buyBody struct {
	UID 		uint 	`json:"uid"`
	Count		uint	`json:"count"`
	Payment		string	`json:"payment"`
}

// Buy
// @Summary Buy product
// @Description buy product in shop by ID
// @ID buy-product-in-shop
// @Accept	json
// @Produce  json
// @Param   id      path   int     true  "ID of the shop"
// @Param   pid      path   int     true  "ID of the product"
// @Param   purchase_info body buyBody true "UID count and type of payment"
// @Success 200
// @Failure 404 {string} string "Shop don't find"
// @Failure 404 {string} string "Product don't find"
// @Failure 400 {string} string "Incorrect body"
// @Failure 400 {string} string "Count is null"
// @Failure 400 {string} string "NegCount"
// @Failure 502 {string} string "Faield to connect service purchases"
// @Failure 404 {string} string "Body is null"
// @Router /shops/{id}/{pid} [put]
func (sp *ShopHandler) Buy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := getShopID(vars)
	productID := getProductID(vars)

	p := &pm.Purchase{}

	d := json.NewDecoder(r.Body)
	if err := d.Decode(p); err == io.EOF {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Body is null")
		return
	} else if err != nil {
		logAndWriteAboutInternalError(w, err, "Buy")
		return
	}

	if p.UID == 0 || p.Count == 0 || p.Payment == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Incorrect body")
		return
	}
	
	p.ShopID = id
	p.ProductID = productID

	if err := sp.buy(p); err != nil {
		if err, ok := validateErrorNotFound(err); ok {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		} else if err, ok := validateErrorBadRequest(err); ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		} else if err == ErrConnectToPurchases { 
			w.WriteHeader(http.StatusBadGateway)
			w.Write([]byte(err.Error()))
			return
		} else {
			logAndWriteAboutInternalError(w, err, "Buy")
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

// func (sh *ShopHandler) GetPurchases(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	uid := getUID(vars)

// 	ps, err := sh.getAllPurchases(uid)
// 	if err == pc.ErrNotFound {
// 		w.WriteHeader(http.StatusNotFound)
// 		w.Write([]byte(err.Error()))
// 		return
// 	} else if err != nil {
// 		logAndWriteAboutInternalError(w, err, "GetPurchases")
// 		return
// 	}

// 	data, err := json.Marshal(ps)
// 	if err != nil {
// 		logAndWriteAboutInternalError(w, err, "GetPurchases")
// 		return
// 	}

// 	w.Write(data)
// 	w.WriteHeader(http.StatusOK)
// }

// GetAll
// @Summary GetAll shops
// @Description get all shops
// @ID get-all-shops
// @Produce  json
// @Success 200 {array} m.Shop
// @Failure 404 {string} string "Shop don't find"
// @Router /shops/ [get]
func (sh *ShopHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	shops, err := sh.c.GetAll()

	if err == pc.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	} else if err != nil {
		logAndWriteAboutInternalError(w, err, "GetAll")
		return
	}

	data, err := json.Marshal(shops)
	if err != nil {
		logAndWriteAboutInternalError(w, err, "GetAll")
		return
	}

	w.Write(data)
}

// AddProduct
// @Summary Add a product to the shop
// @Description Add product by id to the shop by id
// @ID add-product-to-shop-by-id
// @Produce  json
// @Param   token header string true "Authirization Header with hashed secret_key looks like: Token fhdjfho23h4ore"
// @Param   id      path   int     true  "ID of the shop"
// @Param   pid      path   int     true  "ID of the product"
// @Success 200
// @Failure 404 {string} string "Shop don't find"
// @Failure 404 {string} string "Product don't find"
// @Failure 403
// @Router /shops/{id}/{pid} [post]
func (sh *ShopHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := getShopID(vars)
	productID := getProductID(vars)

	shop, err := sh.c.Get(id)
	if err == sc.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	} else if err != nil {
		logAndWriteAboutInternalError(w, err, "AddProduct")
		return
	}

	shop.ShopProducts = append(shop.ShopProducts, m.ShopProduct{ProductID: productID})
	if err := sh.c.Update(shop); err == product.ErrNotFound {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	} else if err != nil {
		logAndWriteAboutInternalError(w, err, "AddProduct")
		return
	}
}

// AddCount
// @Summary Add a count of the product
// @Description Add count of the product
// @ID add-count-of-product-to-shop-by-id
// @Produce  json
// @Param   token header string true "Authirization Header with hashed secret_key looks like: Token fhdjfho23h4ore"
// @Param   id      path   int     true  "ID of the shop"
// @Param   pid      path   int     true  "ID of the product"
// @Param   count      path   int     true  "ID of the product"
// @Success 200
// @Failure 400 {string} string "Product don't find"
// @Router /shops/{id}/{pid}/{count} [put]
func (sh *ShopHandler) AddCount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := getShopID(vars)
	productID := getProductID(vars)
	count := getCount(vars)

	if err := sh.c.AddCount(id, productID, count); err == shop.ErrProductNotFound {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	} else if err != nil {
		logAndWriteAboutInternalError(w, err, "AddCount")
		return
	}

}


// CreateShop
// @Summary Create shop
// @Description Create a shop
// @ID create-shop
// @Accept json
// @Produce  json
// @Param   shop_info body m.ShopInfo true "Information about shop"
// @Success 200
// @Failure 400 {string} string "Shop exist"
// @Failure 400 {string} string "Incorrect body"
// @Failure 400 {string} string "body is null"
// @Router /shops/ [post]
func (sh *ShopHandler) CreateShop(w http.ResponseWriter, r *http.Request) {
	s := m.ShopInfo{}

	d := json.NewDecoder(r.Body)

	if err := d.Decode(&s); err == io.EOF{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Body is null")
		return
	} else if err != nil {
		logAndWriteAboutInternalError(w, err, "CreateShop")
		return
	}

	if s.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Incorrect body")
		return
	}

	shop := &m.Shop{ShopInfo: s}
	if err := sh.c.Create(shop); err == sc.ErrExist {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	} else if err != nil {
		logAndWriteAboutInternalError(w, err, "CreateShop")
		return
	}

	data, err := json.Marshal(shop.ID)
	if err != nil {
		logAndWriteAboutInternalError(w, err, "CreateShop")
		return
	}

	w.Write(data)
}

func getShopID(v map[string]string) uint {
	return getUINT(v, "id")
}

func getProductID(v map[string]string) uint {
	return getUINT(v, "pid")
}

func getUID(v map[string]string) uint {
	return getUINT(v, "uid")
}

func getCount(v map[string]string) uint {
	return getUINT(v, "count")
}

func getUINT(v map[string]string, name string) uint {
	get := v[name]

	_uint, _ := strconv.ParseUint(get, 10, 64)

	return uint(_uint)
}

// TODO check for is shop not found or product

func (sp *ShopHandler) buy(p *pm.Purchase) error {
	url := fmt.Sprintf("%s/purchases/%v", sp.PurhacesNetwork, p.UID)
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}

	body := bytes.NewReader(data)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}

	setAuthHeader(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if opError, ok := err.(*u.Error).Err.(*net.OpError); 
		ok && opError.Op == "dial" && opError.Err != nil {
			return ErrConnectToPurchases
		} else {
			return err
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		if err := checkError(resp); err != nil {
			return err
		}
	} else if resp.StatusCode == http.StatusBadRequest {
		if err := checkError(resp); err != nil {
			return err
		}
	} else if resp.StatusCode != http.StatusCreated {
		return unexcpetedCode(resp.StatusCode)
	}
	
	if err := sp.c.SubCount(p.ShopID, p.ProductID, p.Count); err != nil {
		return err
	}

	return nil
}

func (sp *ShopHandler) getAllPurchases(UID uint) ( []pm.Purchase, error ) {
	url := fmt.Sprintf("%s/purchases/%v", sp.PurhacesNetwork, UID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	setAuthHeader(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, pc.ErrNotFound
	} else if resp.StatusCode != http.StatusOK {
		return nil, unexcpetedCode(resp.StatusCode)
	}

	var ps []pm.Purchase

	d := json.NewDecoder(resp.Body)
	if err := d.Decode(&ps); err != nil{
		return nil, err
	}

	return ps, nil
}

func setAuthHeader(r *http.Request) {
	sha := sha512.New()
	sha.Write([]byte(middlewares.SecretKey))

	data := sha.Sum(nil)

	token := fmt.Sprintf("Token %s", hex.EncodeToString(data))

	r.Header.Add("Authorization", token)
}

func unexcpetedCode(code int) error {
	return fmt.Errorf("Unexpected code: %v", code)
}

func checkError(resp *http.Response) (error) {
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(data)

	if buf.Len() != 0 {
		return errors.New(buf.String())
	}

	return nil
}

func validateErrorNotFound(err error) (error, bool) {
	if err.Error() == pc.ErrInvalidShopID.Error() {
		return pc.ErrInvalidShopID, true
	} else if err.Error() == pc.ErrInvalidProductID.Error() {
		return pc.ErrInvalidProductID, true
	}

	return err, false
}

func validateErrorBadRequest(err error) (error, bool) {
	if err.Error() == pc.ErrCountNull.Error() {
		return pc.ErrCountNull, true
	} else if err.Error() == sc.ErrNegCount.Error() {
		return sc.ErrNegCount, true
	}

	return err, false
}