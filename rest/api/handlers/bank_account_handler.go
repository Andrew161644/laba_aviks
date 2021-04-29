package handlers

import (
	"fmt"
	"github.com/Andrew161644/avicks_laba/api/database/models"
	"github.com/Andrew161644/avicks_laba/api/handlers/views"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (app *Injection) BankAccountHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id, name, _ := app.UserSession.GetCurrentUserIdName(r)
		var curr string
		var bankAccValue string
		var accId string
		var bankAccs, _ = app.DataBase.GetAllBankAccountsByUserId(models.BankAccount{UserId: id})
		var currencies, _ = app.DataBase.GetAllCurrencies()
		var bankAccViewModel []views.BankAccount
		tmpl, _ := template.ParseFiles("../resources/html/bank/bank_account.html")
		var bankId = r.FormValue("ID")

		for _, account := range bankAccs.BankAccounts {
			var curTemp string
			for _, currency := range currencies {
				if currency.ID == account.CurrencyId {
					curTemp = currency.Name
					curr = currency.Name
				}
			}
			if bankId == account.ID {
				bankAccValue = fmt.Sprintf("%.2f", account.Value)
				accId = account.ID
			}
			bankAccViewModel = append(bankAccViewModel, views.BankAccount{
				ID:       account.ID,
				Currency: curTemp,
			})
		}
		err := tmpl.Execute(w, views.ConcreteBankAccount{
			Title:        "BankAccount",
			ID:           accId,
			UserName:     name,
			AccountValue: bankAccValue,
			Currency:     curr,
			BankAccounts: bankAccViewModel,
			Currencies:   currencies,
		})
		if err != nil {
			log.Fatal(err)
		}
	case "POST":
		var add_value = r.FormValue("add_value")
		var ID = r.FormValue("ID")
		var acc, _ = app.DataBase.GetBankAccountById(ID)
		log.Println(acc)
		var val, _ = strconv.ParseFloat(add_value, 64)
		acc.Value += val
		app.DataBase.UpdateBankAccountById(acc)
		log.Println(add_value)
		r.Method = "GET"
		app.BankAccountHandler(w, r)
	}
}

func (app *Injection) BankAccountSecureHandler(w http.ResponseWriter, r *http.Request) {
	var isLogin = app.UserSession.IsUserLogin(r)
	if !isLogin {
		app.LoginDevHandler(w, r)
	} else {
		app.BankAccountHandler(w, r)
	}
}
