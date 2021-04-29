package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func (app *Injection) CompanyHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("../resources/html/company.html")
	err := tmpl.Execute(w, app.AppCreateViewData("Ваша компания", r))
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Error")
	}
}
