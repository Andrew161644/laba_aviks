package models

type BankAccount struct {
	ID         int `json:"id"`
	CurrencyId int `json:"currencyId"`
	UserId     int `json:"userId"`
}

type BankAccountList struct {
	BankAccounts []BankAccount `json:"bank_accounts"`
}
