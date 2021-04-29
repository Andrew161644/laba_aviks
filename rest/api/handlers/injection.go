package handlers

import (
	. "github.com/Andrew161644/avicks_laba/api/config"
	. "github.com/Andrew161644/avicks_laba/api/database/providers"
	. "github.com/Andrew161644/avicks_laba/api/session"
	"log"
	"net/http"
)

// структура для внедрения
type Injection struct {
	DataBase    *Database
	UserSession *UserSession
	Conf        *Config
}

// создание сессии
func (app *Injection) CreateNewSession() {
	var session = CreateNewUserSession()
	app.UserSession = &session
}

type ViewData struct {
	Title    string
	UserName string
}

// Создание модели(динамические поля) для страниц
func (app Injection) AppCreateViewData(title string, r *http.Request) ViewData {
	_, name, err := app.UserSession.GetCurrentUserIdName(r)
	log.Println(name, err)
	if err != nil {
		return ViewData{
			Title:    title,
			UserName: "Гость",
		}
	}
	return ViewData{
		Title:    title,
		UserName: name,
	}
}
