package models

type Currency struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CurrencyList struct {
	Currencies []Currency `json:"currencies"`
}
