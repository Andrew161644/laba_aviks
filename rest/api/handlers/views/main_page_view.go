package views

// Структура хранящая динамические поля для отображения в html
type ViewData struct {
	Title    string
	UserName string
}

// Эти структуры находятся во views
type Todo struct {
	ID    int
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}
