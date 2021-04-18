package models

type Currency struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CurrencyList struct {
	Currencies []Currency `json:"currencies"`
}
