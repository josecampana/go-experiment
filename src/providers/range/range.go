package RangeProvider

import (
	"log"

	"github.com/go-resty/resty/v2"
)

type RangeResponse struct {
	Ids   []string                 `json:"ids"`   // Array de ids
	Items []map[string]interface{} `json:"items"` // Array de items con propiedades dinámicas
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
