package handlers

import (
	"fmt"
	"github.com/Andrew161644/avicks_laba/api/database/models"
	"github.com/Andrew161644/avicks_laba/api/handlers/views"
	. "github.com/Andrew161644/avicks_laba/api/session"
	"html/template"
	"net/http"
)

func (app *Injection) RegistartionHandlerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		data := views.ViewData{
			Title: "Регистрация",
		}
		tmpl, _ := template.ParseFiles("../static/registration.html")
		err := tmpl.Execute(w, data)
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

		var _, err = app.UserSession.CheckIfUserNameExist(UserData{Name: name, Password: password}, app.DataBase)

		if err == nil {
			app.DataBase.AddUser(models.UserModel{
				Name:     name,
				Password: password,
			})
			app.UserSession.LoginUser(w, r, app.DataBase, &UserData{Name: name, Password: password})

			data := views.ViewData{
				Title: "Главная",
			}
			tmpl, _ := template.ParseFiles("../static/main.html")
			err := tmpl.Execute(w, data)
			if err != nil {
				fmt.Println(err)
				fmt.Fprintf(w, "Error")
			}
		} else {
			data := views.ViewData{
				Title: "Регистрация",
			}
			tmpl, _ := template.ParseFiles("../static/registration.html")
			err := tmpl.Execute(w, data)
			if err != nil {
				fmt.Println(err)
				fmt.Fprintf(w, "Error")
			}
		}
	}
}
