package handlers

import (
	. "github.com/Andrew161644/avicks_laba/api/handlers/views"
	"html/template"
	"net/http"
)

// это находится в папке handler
func (app *Injection) HandlerExample(w http.ResponseWriter, r *http.Request) {
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{ID: 1, Title: "Task 1", Done: false},
			{ID: 2, Title: "Task 2", Done: true},
			{ID: 3, Title: "Task 3", Done: true},
		},
	}
	tmpl := template.Must(template.ParseFiles("../resources/html/example.html"))
	tmpl.Execute(w, data)
}
