package routes

import "net/http"
import "github.com/Andrew161644/avicks_laba/api/handlers"

func AddRoutes(app handlers.Injection) {
	http.HandleFunc("/main_page", app.HelloPageHandler)
	http.HandleFunc("/about", app.AboutHandler)
	http.HandleFunc("/developers", app.DevelopersHandler)
	http.HandleFunc("/news", app.NewsHandler)
	http.HandleFunc("/find_dev", app.FindDevHandler)
}
