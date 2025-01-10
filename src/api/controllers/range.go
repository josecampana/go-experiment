package controllers

import (
	constants "b2b-service-pmp/src/api/constants"
	RangeProvider "b2b-service-pmp/src/providers/range"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(constants.CONTENT_TYPE, constants.JSON)
	params := mux.Vars(r)
	id := params["id"]

	// rangeData := RangeProvider.Get2(id)
	// product := rangeData.Items[0]
	product := RangeProvider.Get(id).Items[0]

	genericProduct := product["genericProduct"]
	fakeid := product["_id"]
	color := product["colours"]
	product["_fakeId"] = product["_id"].(string) // this is a type assertion
	product["___fistro"] = "fistro"

	//to avoid the error of unused vars:
	log.Println(fakeid, product["_fakeId"], color, genericProduct)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
