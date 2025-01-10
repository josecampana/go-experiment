package RangeProvider

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type RangeResponse struct {
	Ids   []string                 `json:"ids"`   // Array de ids
	Items []map[string]interface{} `json:"items"` // Array de items con propiedades dinámicas
}

func GetNativo(id string) RangeResponse {
	url := "https://services.ifb.ingka.com/range/v3/us/en/products?productContent=PRICE%2CPRICES%2CAVAILABILITY%2CVARIANTS%2CEXPAND_CHILDS%2CPARSE_IMAGES%2CPACKAGE_INFO%2CSAFETY_WARNING%2CCATEGORIES&ids=" + id

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("request error calling %s: %v", url, err)
	}
	defer resp.Body.Close() // Cerrar el cuerpo de la respuesta al final

	// Leer el cuerpo de la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading response body: %v", err)
	}

	// var apiResponse interface{} //generic
	var apiResponse RangeResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		log.Fatalf("error deserializing JSON response: %v", err)
	}
	return apiResponse
}

func Get(id string) RangeResponse {
	client := resty.New()
	url := "https://services.ifb.ingka.com/range/v3/us/en/products?productContent=PRICE%2CPRICES%2CAVAILABILITY%2CVARIANTS%2CEXPAND_CHILDS%2CPARSE_IMAGES%2CPACKAGE_INFO%2CSAFETY_WARNING%2CCATEGORIES&ids=" + id

	resp, err := client.R().
		SetResult(&RangeResponse{}).
		Get(url)

	if err != nil {
		log.Fatalf("request error calling %s: %v", url, err)
	}

	response := resp.Result().(*RangeResponse)

	return *response
}
