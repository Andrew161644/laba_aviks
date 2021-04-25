package models

type BankAccount struct {
	ID         string `json:"id"`
	Value      int    `json:"value"`
	CurrencyId int    `json:"currencyId"`
	UserId     int    `json:"userId"`
}

type BankAccountList struct {
	BankAccounts []BankAccount `json:"bank_accounts"`
}
