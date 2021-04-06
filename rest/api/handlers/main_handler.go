package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Вернуть главную страницу
func (app *Injection) HelloPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("../resources/html/main.html")
	err := tmpl.Execute(w, app.AppCreateViewData("Главная", r))
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Error")
	}
}

// Вернуть страницу О нас
func (app *Injection) AboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("../resources/html/about.html")
	err := tmpl.Execute(w, app.AppCreateViewData("О нас", r))
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Error")
	}
}

// Вернуть страницу Разработчики
func (app *Injection) DevelopersHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("../resources/html/members.html")
	err := tmpl.Execute(w, app.AppCreateViewData("Разработчики", r))
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Error")
	}
}

// Вернуть страницу новости
func (app *Injection) NewsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("../resources/html/news.html")
	err := tmpl.Execute(w, app.AppCreateViewData("Новости", r))
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Error")
	}
}

func (app *Injection) DevelopersSecureHandler(w http.ResponseWriter, r *http.Request) {
	var isLogin = app.UserSession.IsUserLogin(r)
	if !isLogin {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	id, _, error := app.UserSession.GetCurrentUserIdName(r)
	if error != nil {
		log.Fatal(error)
	}
	log.Println(id)
	user, err := app.DataBase.GetUserById(id)
	log.Println(user)
	if err != nil {
		log.Println(error)
	}
	if user.RoleId != 1 {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	app.DevelopersHandler(w, r)
}
