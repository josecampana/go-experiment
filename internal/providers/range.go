package RangeProvider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	config "github.com/ingka-group-digital/b2b-service-pmp/configs"
	ContextHelper "github.com/ingka-group-digital/b2b-service-pmp/internal/context"

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
	Context    context.Context
}

func Get(ids [](string), options GetOptions) (RangeResponse, error) {
	logger := ContextHelper.Logger(options.Context)
	transactionId := ContextHelper.TransactionId(options.Context)

	urlBase := config.Get().Providers.Range
	client := resty.New()

	URL := fmt.Sprintf("%s/v3/%s/%s/products?productContent=%s&ids=%s", urlBase, options.RetailUnit, options.Language, url.QueryEscape(options.Content), url.QueryEscape(strings.Join(ids, ",")))
	if options.PostalCode != "" {
		URL += "&postalCode=" + options.PostalCode
	}
	if options.Store != "" {
		URL += "&store=" + options.Store
	}

	logger.Debug("calling external API", "url", URL, "method", "GET")

	resp, err := client.R().
		SetContext(options.Context).
		SetHeader("Content-Type", "application/json").
		SetHeader("x-transaction-id", transactionId).
		SetResult(&RangeResponse{}).
		Get(URL)

	if err != nil {
		logger.Error("request error calling external API", err)

		return RangeResponse{}, err
	}

	response := resp.Result().(*RangeResponse)

	return *response, nil
}

//logs:
//https://www.youtube.com/watch?v=gd_Vyb5vEw0
