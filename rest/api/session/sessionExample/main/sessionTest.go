package main

import (
	"fmt"
	"github.com/Andrew161644/avicks_laba/api/database/providers"
	"github.com/Andrew161644/avicks_laba/api/session"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"net/http"
)

var userSession = session.UserSession{
	Store: sessions.NewCookieStore(securecookie.GenerateRandomKey(32)),
}
var db, _ = providers.Connect("localhost", 5432, "postgres", "postgres", "postgres")

func secret(w http.ResponseWriter, r *http.Request) {

	var userLogin = userSession.IsUserLogin(r)
	if !userLogin {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	fmt.Fprintln(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	userSession.LoginUser(w, r, db, &session.UserData{Name: "TestUser", Password: "test1"})
}

func logout(w http.ResponseWriter, r *http.Request) {
	userSession.LogoutUser(w, r)
}

func main() {
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(":8080", nil)
}
