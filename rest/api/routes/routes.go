package routes

import (
	. "github.com/Andrew161644/avicks_laba/api/handlers"
	"net/http"
)

func AddRoutes(app Injection) {
	http.HandleFunc("/main_page", app.HelloPageHandler)
	http.HandleFunc("/about", app.AboutHandler)
	http.HandleFunc("/developers", app.DevelopersHandler)
	http.HandleFunc("/news", app.NewsHandler)
	http.HandleFunc("/find_dev", app.FindDevHandler)
	http.HandleFunc("/login", app.LoginDevHandler)
	http.HandleFunc("/registration", app.RegistartionHandlerHandler)
}
