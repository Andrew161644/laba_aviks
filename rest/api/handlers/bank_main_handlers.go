package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func (app *Injection) BankMainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("../resources/html/bank_main.html")
	err := tmpl.Execute(w, app.AppCreateViewData("Ваши счета", r))
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Error")
	}
}
