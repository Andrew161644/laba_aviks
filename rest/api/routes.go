package main

import (
	"github.com/Andrew161644/avicks_laba/api/handlers"
	"net/http"
)

// пути для приложения(сопоставляем обработчикам)
func AddRoutes(app handlers.Injection) {
	http.HandleFunc("/main_page", app.HelloPageHandler)
	http.HandleFunc("/about", app.AboutHandler)
	http.HandleFunc("/developers", app.DevelopersSecureHandler)
	http.HandleFunc("/news", app.NewsHandler)
	http.HandleFunc("/login", app.LoginDevHandler)
	http.HandleFunc("/registration", app.RegistartionHandlerHandler)
	http.HandleFunc("/logout", app.LogOutHandler)
	http.HandleFunc("/find_dev", app.ResumeSecureHandler)
	http.HandleFunc("/call_rabbit", app.SendTaskExample)
	http.HandleFunc("/call_python_test", app.NeiroServiceTest)
	http.HandleFunc("/bank_account", app.BankAccountHandler)
	http.HandleFunc("/bank_main", app.BankMainHandler)
	http.HandleFunc("/example", app.HandlerExample)
	http.HandleFunc("/company", app.CompanyHandler)
}
