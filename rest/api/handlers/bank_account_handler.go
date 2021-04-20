package handlers

import (
	"fmt"
	"github.com/Andrew161644/avicks_laba/api/handlers/views"
	"html/template"
	"net/http"
)

func (app *Injection) BankAccountHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("../resources/html/bank_account.html")
	err := tmpl.Execute(w, views.ViewData{Title: "dfds", UserName: "Fidan"})
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Error")
	}
}
