package handlers

import "net/http"

// Обработчик выхода из аккаунта
func (app *Injection) LogOutHandler(w http.ResponseWriter, r *http.Request) {
	app.UserSession.LogoutUser(w, r)
	app.HelloPageHandler(w, r)
}
