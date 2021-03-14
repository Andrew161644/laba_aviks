package session

import (
	"errors"
	. "github.com/Andrew161644/avicks_laba/api/database/providers"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

type IUserSession interface {
	LoginUser(w http.ResponseWriter, r *http.Request, dataBase *Database, data *UserData)
	LogoutUser(w http.ResponseWriter, r *http.Request)
	IsUserLogin(r *http.Request) bool
	GetCurrentUserId(r *http.Request) (int, error)
	CheckIfUserNamePasswordExist(data UserData, database *Database) (int, error)
}

var notRegister = errors.New("math: square root of negative number")

type UserSession struct {
	Store *sessions.CookieStore
}

func CreateNewUserSession() UserSession {
	return UserSession{
		Store: sessions.NewCookieStore(securecookie.GenerateRandomKey(32)),
	}
}

func (userSession *UserSession) LoginUser(w http.ResponseWriter, r *http.Request, dataBase *Database, data *UserData) {
	session, _ := userSession.Store.Get(r, "cookie-name")

	var id, err = userSession.CheckIfUserNamePasswordExist(*data, dataBase)

	if err == nil {
		session.Values["authenticated"] = true
		session.Values["id"] = id
		log.Println(session.Values["id"].(int))
	}
	session.Save(r, w)
}

func (userSession *UserSession) LogoutUser(w http.ResponseWriter, r *http.Request) {
	session, _ := userSession.Store.Get(r, "cookie-name")

	session.Values["authenticated"] = false
	session.Values["id"] = nil

	session.Save(r, w)
}

func (userSession *UserSession) IsUserLogin(r *http.Request) bool {
	session, _ := userSession.Store.Get(r, "cookie-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return false
	}
	return true
}

func (userSession *UserSession) GetCurrentUserId(r *http.Request) (int, error) {
	if userSession.IsUserLogin(r) {
		session, _ := userSession.Store.Get(r, "cookie-name")
		return session.Values["id"].(int), nil
	}
	return 0, notRegister
}

func (userSession *UserSession) CheckIfUserNamePasswordExist(data UserData, database *Database) (int, error) {
	var user, err = database.GetUserByNameAndPassword(data.Name, data.Password)
	if err != nil {
		return 0, err
	} else {
		return user.ID, nil
	}
}
