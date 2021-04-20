package currency_exchange_client

import (
	"encoding/json"
	"log"
	"net/http"
)

var ExchangeUri = "http://neiro_service:5000/currency"

func CallExchangeCurrency(uri string, rm RequestCurrencyExchangeModel) (*ResponseCurrencyExchangeModel, error) {
	data := rm.AsUrlValues()
	resp, err := http.PostForm(uri, data)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var res ResponseCurrencyExchangeModel

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &res, nil
}
