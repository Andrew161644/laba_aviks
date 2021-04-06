package models

type Index struct {
	Id          int     `json:"id"`
	PastYear    int     `json:"pastYear"`
	CurrentYear int     `json:"currentYear"`
	Value       float64 `json:"value"`
	CurrencyId  int     `json:"currencyId"`
}
