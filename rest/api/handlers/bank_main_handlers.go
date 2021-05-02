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

func (app *Injection) BankMainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl, _ := template.ParseFiles("../resources/html/bank/bank_main.html")
		id, name, _ := app.UserSession.GetCurrentUserIdName(r)
		var bankAccs, err = app.DataBase.GetAllBankAccountsByUserId(models.BankAccount{UserId: id})
		var currencies, _ = app.DataBase.GetAllCurrencies()
		var bankAccViewModel []views.BankAccount

		for _, account := range bankAccs.BankAccounts {
			var cur string
			for _, currency := range currencies {
				if currency.ID == account.CurrencyId {
					cur = currency.Name
					break
				}
			}
			bankAccViewModel = append(bankAccViewModel, views.BankAccount{
				ID:       account.ID,
				Currency: cur,
			})
		}

		err = tmpl.Execute(w, views.BankMainView{
			Title:        "BankMain",
			UserName:     name,
			BankAccounts: bankAccViewModel,
			Currencies:   currencies,
		})

		if err != nil {
			fmt.Println(err)
			fmt.Fprintf(w, "Error")
		}

	case "POST":
		userID, _, _ := app.UserSession.GetCurrentUserIdName(r)
		id, _ := strconv.Atoi(r.FormValue("ID"))
		log.Println("Id :", id)
		log.Println("POST call")
		var bankAcc = models.BankAccount{
			Value:      0,
			CurrencyId: id,
			UserId:     userID,
		}
		var res, err = app.DataBase.AddBankAccount(bankAcc)
		log.Println(err)
		log.Println(res)
		r.Method = "GET"
		app.BankMainHandler(w, r)
	}

}

func (app *Injection) BankMainSecureHandler(w http.ResponseWriter, r *http.Request) {
	var isLogin = app.UserSession.IsUserLogin(r)
	if !isLogin {
		app.LoginDevHandler(w, r)
	} else {
		app.BankMainHandler(w, r)
	}
}
