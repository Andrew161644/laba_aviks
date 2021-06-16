package handlers

import (
	. "github.com/Andrew161644/avicks_laba/api/clients/grpc/client"
	. "github.com/Andrew161644/avicks_laba/api/clients/grpc/task"
	. "github.com/Andrew161644/avicks_laba/api/handlers/converters"
	"log"
	"net/http"
)

func (app *Injection) ExchangePostHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var oldCur string
		var newCur string

		var newBillId = r.FormValue("newBill")
		var oldBillId = r.FormValue("oldBill")
		var value = r.FormValue("value")
		log.Println("Value is ", value)

		var newBill, err = app.DataBase.GetBankAccountById(newBillId)
		if err != nil {
			log.Fatal(err)
		}

		var oldBill, err1 = app.DataBase.GetBankAccountById(oldBillId)
		if err1 != nil {
			log.Fatal(err)
		}
		log.Println(oldBill)
		log.Println(newBill)
		var currencies, _ = app.DataBase.GetAllCurrencies()

		for _, currency := range currencies {
			if currency.ID == newBill.CurrencyId {
				newCur = currency.Name
			}
			if currency.ID == oldBill.CurrencyId {
				oldCur = currency.Name
			}
		}

		var res, err3 = ExchangerRPC(CurrencyExchangeTask{
			Value:               StringToFloat(value),
			CurrentCurrencyName: oldCur,
			NewCurrencyName:     newCur,
		}, app.Conf.RabbitHost)
		if err3 != nil {
			log.Fatal(err)
		}

		newBill.Value += res.Result
		oldBill.Value -= StringToFloat(value)

		app.DataBase.UpdateBankAccountById(newBill)
		app.DataBase.UpdateBankAccountById(oldBill)

		r.Method = "GET"
		r.Form.Del("ID")
		r.Form.Set("ID", oldBill.ID)
		app.BankAccountHandler(w, r)
	}
}
