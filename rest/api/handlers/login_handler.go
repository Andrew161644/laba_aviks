package handlers

import (
	"fmt"
	"github.com/Andrew161644/avicks_laba/api/handlers/views"
	. "github.com/Andrew161644/avicks_laba/api/session"
	"html/template"
	"net/http"
)

func (app *Injection) LoginDevHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		data := views.ViewData{
			Title: "Регистрация",
		}
		tmpl, _ := template.ParseFiles("../static/login.html")
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

		var user, err = app.DataBase.GetUserByNameAndPassword(name, password)
		if err != nil {
			data := views.ViewData{
				Title: "Логин",
			}
			tmpl, _ := template.ParseFiles("../static/login.html")
			err = tmpl.Execute(w, data)
			return
		}
		var userData = ToUserData(user)
		app.UserSession.LoginUser(w, r, app.DataBase, &userData)

		data := views.ViewData{
			Title: "Главная",
		}
		tmpl, _ := template.ParseFiles("../static/main.html")
		err = tmpl.Execute(w, data)
		if err != nil {
			fmt.Println(err)
			fmt.Fprintf(w, "Error")
		}
	}
}
