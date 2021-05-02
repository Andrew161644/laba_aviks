package views

import "github.com/Andrew161644/avicks_laba/api/database/models"

// Структура хранящая динамические поля для отображения в html
type BankMainView struct {
	Title        string
	UserName     string
	BankAccounts []BankAccount
	Currencies   []models.Currency
}
