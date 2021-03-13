package handlers

import (
	"fmt"
	"github.com/Andrew161644/avicks_laba/api/handlers/views"
	"html/template"
	"log"
	"net/http"
)

func (app *Injection) HelloPageHandler(w http.ResponseWriter, r *http.Request) {
	data := views.ViewData{
		Title: "Проект",
	}
	tmpl, _ := template.ParseFiles("../static/main.html")
	err := tmpl.Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, "Error")
	}
}

func getAdmins(app *Injection) []string {
	admins, err := app.DataBase.GetUserByRole(1)
	if err != nil {
		log.Fatal(err)
	}
	var res []string
	for _, user := range admins.Users {
		res = append(res, user.Name)
	}
	return res
}
