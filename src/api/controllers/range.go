package controllers

import (
	constants "b2b-service-pmp/src/api/constants"
	RangeProvider "b2b-service-pmp/src/providers/range"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(constants.CONTENT_TYPE, constants.JSON)
	params := mux.Vars(r)
	retailUnit := params["retailUnit"]
	language := params["language"]

	queryParams := r.URL.Query()
	ids := strings.Split(queryParams.Get("ids"), ",")

	productContent := queryParams.Get("productContent")
	if productContent == "" {
		productContent = "PRICE,PRICES,AVAILABILITY,VARIANTS,EXPAND_CHILDS,PARSE_IMAGES,PACKAGE_INFO"
	}

	postalCode := queryParams.Get("postalCode")
	store := queryParams.Get("store")

	options := RangeProvider.GetOptions{
		RetailUnit: retailUnit,
		Language:   language,
		PostalCode: postalCode,
		Store:      store,
		Content:    productContent,
	}

	products := RangeProvider.Get(ids, options)

	// genericProduct := product["genericProduct"]
	// fakeid := product["_id"]
	// color := product["colours"]
	// product["_fakeId"] = product["_id"].(string) // this is a type assertion
	// product["___fistro"] = "fistro"

	// //to avoid the error of unused vars:
	// log.Println(fakeid, product["_fakeId"], color, genericProduct)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
