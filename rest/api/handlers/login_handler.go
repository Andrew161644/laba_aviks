package handlers

import (
	"fmt"
	. "github.com/Andrew161644/avicks_laba/api/session"
	"html/template"
	"net/http"
)

// обработчик login
func (app *Injection) LoginDevHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl, _ := template.ParseFiles("../resources/html/login.html")
		err := tmpl.Execute(w, app.AppCreateViewData("Логин", r))

		if err != nil {
			fmt.Println(err)
			fmt.Fprintf(w, "Error")
		}

	case "POST":

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		name := r.FormValue("name")
		password := r.FormValue("password")

		var user, err = app.DataBase.GetUserByNameAndPassword(name, password)
		if err != nil {
			tmpl, _ := template.ParseFiles("../resources/html/login.html")
			err = tmpl.Execute(w, app.AppCreateViewData("Логин", r))
			return
		}
		var userData = ToUserData(user)
		app.UserSession.LoginUser(w, r, app.DataBase, &userData)

		r.Method = "GET"
		app.BankMainHandler(w, r)
	}
}
