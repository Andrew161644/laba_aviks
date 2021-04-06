package handlers

import (
	"fmt"
	"github.com/Andrew161644/avicks_laba/api/database/models"
	"html/template"
	"log"
	"net/http"
)

func (app *Injection) ResumeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		{
			log.Println("GET resume handler")
			tmpl, _ := template.ParseFiles("../resources/html/resume.html")
			err := tmpl.Execute(w, app.AppCreateViewData("Резюме", r))
			if err != nil {
				fmt.Println(err)
				fmt.Fprintf(w, "Error")
			}
		}
	case "POST":
		{
			log.Println("POST resume handler")

			// нужно обработать эти поля с помощью провайдера
			name := r.FormValue("name")
			email := r.FormValue("email")
			speciality := r.FormValue("speciality")
			about := r.FormValue("about")

			app.DataBase.AddResume(models.Resume{
				Name:       name,
				Email:      email,
				Speciality: speciality,
				About:      about,
			})

			// возвращает страницу
			tmpl, _ := template.ParseFiles("../resources/html/resume.html")
			err := tmpl.Execute(w, app.AppCreateViewData("Резюме", r))
			if err != nil {
				fmt.Println(err)
				fmt.Fprintf(w, "Error")
			}
		}
	}
}
