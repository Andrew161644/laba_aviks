package handlers

import (
	"fmt"
	"github.com/Andrew161644/avicks_laba/api/handlers/views"
	"html/template"
	"net/http"
)

func (app *Injection) HelloPageHandler(w http.ResponseWriter, r *http.Request) {
	data := views.ViewData{
		Title: "Главная",
	}
	tmpl, _ := template.ParseFiles("../static/main.html")
	err := tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Error")
	}
}
func (app *Injection) AboutHandler(w http.ResponseWriter, r *http.Request) {
	data := views.ViewData{
		Title: "О нас",
	}
	tmpl, _ := template.ParseFiles("../static/about.html")
	err := tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Error")
	}
}

func (app *Injection) DevelopersHandler(w http.ResponseWriter, r *http.Request) {
	data := views.ViewData{
		Title: "Об участниках",
	}
	tmpl, _ := template.ParseFiles("../static/members.html")
	err := tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Error")
	}
}

func (app *Injection) NewsHandler(w http.ResponseWriter, r *http.Request) {
	data := views.ViewData{
		Title: "Новости",
	}
	tmpl, _ := template.ParseFiles("../static/news.html")
	err := tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Error")
	}
}

func (app *Injection) FindDevHandler(w http.ResponseWriter, r *http.Request) {
	data := views.ViewData{
		Title: "Резюме",
	}
	tmpl, _ := template.ParseFiles("../static/resume.html")
	err := tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Error")
	}
}

func (app *Injection) LoginDevHandler(w http.ResponseWriter, r *http.Request) {
	data := views.ViewData{
		Title: "Войти",
	}
	tmpl, _ := template.ParseFiles("../static/login.html")
	err := tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Error")
	}
}

func (app *Injection) RegDevHandler(w http.ResponseWriter, r *http.Request) {
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

//func getAdmins(app *Injection) []string {
//	admins, err := app.DataBase.GetUserByRole(1)
//	if err != nil {
//		log.Fatal(err)
//	}
//	var res []string
//	for _, user := range admins.Users {
//		res = append(res, user.Name)
//	}
//	return res
//}
