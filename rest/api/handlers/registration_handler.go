package handlers

import (
	"fmt"
	"github.com/Andrew161644/avicks_laba/api/database/models"
	. "github.com/Andrew161644/avicks_laba/api/session"
	"html/template"
	"net/http"
)

func (app *Injection) RegistartionHandlerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl, _ := template.ParseFiles("../resources/html/registration.html")
		err := tmpl.Execute(w, app.AppCreateViewData("Регистрация", r))
		if err != nil {
			fmt.Println(err)
			fmt.Fprintf(w, "Error")
		}
	case "POST":

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		var userData = UserData{
			Name:     r.FormValue("name"),
			Password: r.FormValue("password"),
		}

		var _, isExist = app.UserSession.CheckIfUserNameExist(userData, app.DataBase)

		if isExist == false && IsValid(userData) {
			app.DataBase.AddUser(models.UserModel{
				Name:     userData.Name,
				Password: userData.Password,
			})
			app.UserSession.LoginUser(w, r, app.DataBase, &userData)
			app.HelloPageHandler(w, r)

		} else {
			tmpl, _ := template.ParseFiles("../resources/html/registration.html")
			err := tmpl.Execute(w, app.AppCreateViewData("Регистрация", r))
			if err != nil {
				fmt.Println(err)
				fmt.Fprintf(w, "Error")
			}
		}
	}
}
