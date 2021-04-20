package currency_exchange_client

type ResponseCurrencyExchangeModel struct {
	Value               float64 `json:"value"`
	CurrentCurrencyName string  `json:"currentCurrencyName"`
}
