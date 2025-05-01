package exchange

import (
	"exchange-rate/intrenal/currency"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
)

const baseUrl = "http://api.exchangerate.host"

type Client struct {
	apiKey  string
	baseUrl string
	client  *resty.Client
}

func New(apiKey string) *Client {
	return &Client{
		apiKey:  apiKey,
		baseUrl: baseUrl,
		client:  resty.New(),
	}
}

//func (exchange *Client) GetQuotes(source currency.Currency) (GetRatesResponse, error) {
//	resp, err := http.Get(fmt.Sprintf("%s/live?access_key=%s&format=1&source=%s", baseUrl, exchange.apiKey, source))
//	if err != nil {
//		return GetRatesResponse{}, err
//	}
//	defer resp.Body.Close()
//
//	body, err := io.ReadAll(resp.Body)
//	if err != nil {
//		return GetRatesResponse{}, err
//	}
//
//	var result GetRatesResponse
//	err = json.Unmarshal(body, &result)
//	if err != nil {
//		return GetRatesResponse{}, err
//	}
//
//	return result, nil
//}

func (ex *Client) GetQuotes(source currency.Currency) (GetRatesResponse, error) {
	queryParams := map[string]string{
		"access_key": ex.apiKey,
		"format":     "1", // json
		"source":     string(source),
	}

	var result GetRatesResponse

	resp, err := ex.client.R().
		SetQueryParams(queryParams).
		SetResult(&result).
		Get(fmt.Sprintf("%s/live", baseUrl))

	if err != nil || resp.StatusCode() != http.StatusOK {
		return GetRatesResponse{}, err
	}

	return result, nil
}
