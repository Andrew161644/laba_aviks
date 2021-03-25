package handlers

import (
	. "github.com/Andrew161644/avicks_laba/api/database/providers"
	"github.com/Andrew161644/avicks_laba/api/handlers/views"
	. "github.com/Andrew161644/avicks_laba/api/session"
	"log"
	"net/http"
)

// структура для внедрения
type Injection struct {
	DataBase    *Database
	UserSession *UserSession
}

// создание сессии
func (app *Injection) CreateNewSession() {
	var session = CreateNewUserSession()
	app.UserSession = &session
}

// Создание модели(динамические поля) для страниц
func (app Injection) AppCreateViewData(title string, r *http.Request) views.ViewData {
	_, name, err := app.UserSession.GetCurrentUserIdName(r)
	log.Println(name, err)
	if err != nil {
		return views.ViewData{
			Title:    title,
			UserName: "Гость",
		}
	}
	return views.ViewData{
		Title:    title,
		UserName: name,
	}
}
