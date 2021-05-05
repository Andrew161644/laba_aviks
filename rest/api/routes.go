package main

import (
	"net/http"

	"github.com/Andrew161644/avicks_laba/api/handlers"
)

// пути для приложения(сопоставляем обработчикам)
func AddRoutes(app handlers.Injection) {

	// call_tests
	http.HandleFunc("/call_rabbit", app.SendTaskExample)
	http.HandleFunc("/call_python_test", app.NeiroServiceTest)

	// bank app
	http.HandleFunc("/bank_main", app.BankMainSecureHandler)
	http.HandleFunc("/bank_account", app.BankAccountSecureHandler)
	http.HandleFunc("/company", app.MyOrganizationSecureHandler)
	http.HandleFunc("/transfer", app.ExchangePostHandler)
	http.HandleFunc("/registration", app.RegistartionHandlerHandler)
	http.HandleFunc("/login", app.LoginDevHandler)
	http.HandleFunc("/logout", app.LogOutHandler)
}
