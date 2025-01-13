package controllers

//just playing a bit with controllers, params and so on...

import (
	"encoding/json"
	"net/http"

	constants "github.com/ingka-group-digital/b2b-service-pmp/internal"

	"github.com/gorilla/mux"
)

type Store struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetStores(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(constants.CONTENT_TYPE, constants.JSON)
	stores := []Store{
		{ID: "406", Name: "Ikea Hospitalet"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stores)
}

func GetStoreById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	queryParams := r.URL.Query()
	storeId := params["storeId"]
	var postalCode *string

	if queryParams.Has("postalCode") {
		val := queryParams.Get("postalCode")
		postalCode = &val
	}

	response := map[string]interface{}{
		"store": storeId,
	}

	if postalCode != nil {
		response["postalCode"] = *postalCode
	}

	w.Header().Set(constants.CONTENT_TYPE, constants.JSON)
	json.NewEncoder(w).Encode(response)
}

func CreateStore(w http.ResponseWriter, r *http.Request) {
	var store Store
	err := json.NewDecoder(r.Body).Decode(&store)

	if err != nil {
		http.Error(w, "error while decoding the body", http.StatusBadRequest)
		return
	}

	//create dynamic response without define them...
	response := map[string]interface{}{
		"message":     "Store created",
		"id":          store.ID,
		"name":        store.Name,
		"i_am_number": 1234565,
		"i_am_float":  123.4565,
		"i_am_bool":   true,
	}

	w.Header().Set(constants.CONTENT_TYPE, constants.JSON)
	json.NewEncoder(w).Encode(response)
}
