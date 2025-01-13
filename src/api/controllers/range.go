package controllers

import (
	constants "b2b-service-pmp/src/api/constants"
	// logger "b2b-service-pmp/src/modules"
	RangeProvider "b2b-service-pmp/src/providers"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"golang.org/x/exp/slog"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(constants.CONTENT_TYPE, constants.JSON)

	params := mux.Vars(r)
	retailUnit := params["retailUnit"]
	language := params["language"]

	queryParams := r.URL.Query()
	rawIds := queryParams.Get("ids")

	if rawIds == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "ids parameter is required"})
		return
	}

	ids := strings.Split(rawIds, ",")

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

	slog.Info("new request for get products", "ids", ids, "options", options)

	products, err := RangeProvider.Get(ids, options)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(products)
}
