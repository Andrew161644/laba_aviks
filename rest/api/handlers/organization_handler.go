package handlers

import (
	. "github.com/Andrew161644/avicks_laba/api/clients/organization_status_client"
	. "github.com/Andrew161644/avicks_laba/api/handlers/converters"
	"github.com/Andrew161644/avicks_laba/api/handlers/views"
	"html/template"
	"log"
	"net/http"
)

func (app *Injection) MyOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		_, name, _ := app.UserSession.GetCurrentUserIdName(r)
		tmpl, _ := template.ParseFiles("../resources/html/bank/company.html")
		err := tmpl.Execute(w, views.MyOrganizationView{
			Title:       "BankAccount",
			UserName:    name,
			Report:      "",
			TitleReport: "Отчет",
			Kk:          "",
			Kn:          "",
			Kfin:        "",
			Kfu:         "",
			Kabsl:       "",
			Kfastl:      "",
			Kcurrl:      "",
		})
		if err != nil {
			log.Fatal(err)
		}
	case "POST":
		_, userName, _ := app.UserSession.GetCurrentUserIdName(r)
		tmpl, _ := template.ParseFiles("../resources/html/bank/company.html")
		/*get form value*/
		var name = r.FormValue("name")
		var borrCap = r.FormValue("borrCap")
		var ownCap = r.FormValue("ownCap")
		var balanceCurr = r.FormValue("balanceCurr")
		var allCash = r.FormValue("allCash")
		var longTermDuties = r.FormValue("longTermDuties")
		var shortTermDuties = r.FormValue("shortTermDuties")
		var shortFinInv = r.FormValue("shortFinInv")
		var shortRec = r.FormValue("shortRec")
		var sumMoney = r.FormValue("sumMoney")

		var data = GetOrgStatusRequestModelUrl(
			name,
			borrCap,
			ownCap,
			balanceCurr,
			allCash,
			longTermDuties,
			shortTermDuties,
			shortFinInv,
			shortRec,
			sumMoney)
		var res, error = CallGetOrgStatusInfoFromData(app.Conf.CalcUri, data)
		if error != nil {
			log.Fatal(error)
		}

		err := tmpl.Execute(w, views.MyOrganizationView{
			Title:       "BankAccount",
			UserName:    userName,
			Report:      res.Report,
			TitleReport: "Отчет",
			Kk:          FloatToString(res.Kk),
			Kn:          FloatToString(res.Kn),
			Kfin:        FloatToString(res.Kfin),
			Kfu:         FloatToString(res.Kfu),
			Kabsl:       FloatToString(res.Kabsl),
			Kfastl:      FloatToString(res.Kfastl),
			Kcurrl:      FloatToString(res.Kcurrl),
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (app *Injection) MyOrganizationSecureHandler(w http.ResponseWriter, r *http.Request) {
	var isLogin = app.UserSession.IsUserLogin(r)
	if !isLogin {
		app.LoginDevHandler(w, r)
	} else {
		app.MyOrganizationHandler(w, r)
	}
}
