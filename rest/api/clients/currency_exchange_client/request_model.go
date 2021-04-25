package currency_exchange_client

import "strconv"

type RequestCurrencyExchangeModel struct {
	Value               float64 `json:"value"`
	CurrentCurrencyName string  `json:"currentCurrencyName"`
	NewCurrencyName     string  `json:"newCurrencyName"`
}

func (obj *RequestCurrencyExchangeModel) AsUrlValues() map[string][]string {
	return map[string][]string{
		"value":               {FloatToString(obj.Value)},
		"currentCurrencyName": {obj.CurrentCurrencyName},
		"newCurrencyName":     {obj.NewCurrencyName},
	}
}

func FloatToString(input_num float64) string {
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}
