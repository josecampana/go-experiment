package RangeProvider

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/go-resty/resty/v2"
	"golang.org/x/exp/slog"
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

func Get(ids [](string), options GetOptions) (RangeResponse, error) {
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
		slog.Error("request error calling external API", "url", URL, "method", "GET", "error", err)

		return RangeResponse{}, fmt.Errorf("error getting range info: %w", err)
	}

	response := resp.Result().(*RangeResponse)

	return *response, nil
}

//logs:
//https://www.youtube.com/watch?v=gd_Vyb5vEw0
