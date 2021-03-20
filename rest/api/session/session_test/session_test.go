package session_test

import (
	"github.com/Andrew161644/avicks_laba/api/database/providers"
	"github.com/Andrew161644/avicks_laba/api/session"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"log"
	"testing"
)

const host = "localhost"

var userSession = session.UserSession{
	Store: sessions.NewCookieStore(securecookie.GenerateRandomKey(32)),
}
var db, err = providers.Connect(host, 5432, "postgres", "postgres", "postgres")

func TestCanCheckUserByNameExist(t *testing.T) {
	id, isExist := userSession.CheckIfUserNameExist(session.UserData{Name: "Fidan"}, db)
	if isExist == false {
		log.Fatal("User not exist")
	}
	log.Println(id)
}

func TestCanCheckUserByNamePasswordExist(t *testing.T) {
	user, isExist := userSession.CheckIfUserNamePasswordExist(session.UserData{Name: "Fidan", Password: "admin"}, db)
	if isExist == false {
		log.Fatal("User not exist")
	}
	log.Println(user)
}
