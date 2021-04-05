package models

type Index struct {
	Id          int `json:"id"`
	PastYear    int `json:"pastYear"`
	CurrentYear int `json:"currentYear"`
	Value       int `json:"value"`
	CurrencyId  int `json:"currencyId"`
}
