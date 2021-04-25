package model

type Configuration struct {
	AMQPConnectionURL string
}

type RequestCurrencyExchangeModel struct {
	Value               float64 `json:"value"`
	CurrentCurrencyName string  `json:"currentCurrencyName"`
	NewCurrencyName     string  `json:"newCurrencyName"`
}

var Config = Configuration{
	AMQPConnectionURL: "amqp://guest:guest@rabbitmq:5672/",
}
