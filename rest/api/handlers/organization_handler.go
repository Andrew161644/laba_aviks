package handlers

import (
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
			TitleReport: "",
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
		_, name, _ := app.UserSession.GetCurrentUserIdName(r)
		tmpl, _ := template.ParseFiles("../resources/html/bank/company.html")
		err := tmpl.Execute(w, views.MyOrganizationView{
			Title:       "BankAccount",
			UserName:    name,
			Report:      "",
			TitleReport: "",
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
