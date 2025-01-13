package RangeProvider

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/go-resty/resty/v2"
)

type RangeResponse struct {
	Ids   []string                 `json:"ids"`   // Array de ids
	Items []map[string]interface{} `json:"items"` // Array de items con propiedades dinámicas
}

type GetOptions struct {
	RetailUnit string
	Language   string
	PostalCode string
	Store      string
	Content    string
}

func Get(ids [](string), options GetOptions) RangeResponse {
	client := resty.New()

	URL := fmt.Sprintf("https://services.ifb.ingka.com/range/v3/%s/%s/products?productContent=%s&ids=%s", options.RetailUnit, options.Language, url.QueryEscape(options.Content), url.QueryEscape(strings.Join(ids, ",")))
	if options.PostalCode != "" {
		URL += "&postalCode=" + options.PostalCode
	}
	if options.Store != "" {
		URL += "&store=" + options.Store
	}

	resp, err := client.R().
		SetResult(&RangeResponse{}).
		Get(URL)

	if err != nil {
		log.Fatalf("request error calling %s: %v", URL, err)
	}

	response := resp.Result().(*RangeResponse)

	return *response
}
