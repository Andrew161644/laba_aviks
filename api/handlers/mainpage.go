package handlers

import (
	"fmt"
	. "github.com/Andrew161644/avicks_laba/api/handlers/views"
	"html/template"
	"net/http"
)

func (app *Application) HelloPageHandler(w http.ResponseWriter, r *http.Request) {
	var user, _ = app.DataBase.GetUserById(1)
	data := ViewData{
		Title:   "Go html",
		Message: user.Name,
	}
	tmpl, _ := template.ParseFiles("/templates/main.html")
	err := tmpl.Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, "Error")
	}
}
