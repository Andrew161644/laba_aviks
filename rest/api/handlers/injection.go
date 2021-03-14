package handlers

import (
	. "github.com/Andrew161644/avicks_laba/api/database/providers"
	. "github.com/Andrew161644/avicks_laba/api/session"
)

type Injection struct {
	DataBase    *Database
	UserSession *UserSession
}

func (app *Injection) CreateNewSession() {
	var session = CreateNewUserSession()
	app.UserSession = &session
}
