package views

import (
	"html/template"
	"net/http"
)

// Структура хранящая динамические поля для отображения в html
type ViewData struct {
	Title    string
	UserName string
}

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl := template.Must(template.ParseFiles("../resources/html/bank_main.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe("localhost:8080/bank_main", nil)
}
