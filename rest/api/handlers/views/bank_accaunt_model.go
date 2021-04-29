package views

import "github.com/Andrew161644/avicks_laba/api/database/models"

type ConcreteBankAccount struct {
	Title        string
	ID           string
	UserName     string
	AccountValue string
	Currency     string
	BankAccounts []BankAccount
	Currencies   []models.Currency
}
