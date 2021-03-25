package session

import (
	"errors"
	"github.com/Andrew161644/avicks_laba/api/database/models"
	. "github.com/Andrew161644/avicks_laba/api/database/providers"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

// Интерфейс для работы с сессией
type IUserSession interface {
	LoginUser(w http.ResponseWriter, r *http.Request, dataBase *Database, data *UserData)
	LogoutUser(w http.ResponseWriter, r *http.Request)
	IsUserLogin(r *http.Request) bool
	GetCurrentUserIdName(r *http.Request) (int, string, error)
	CheckIfUserNameExist(data UserData, database *Database) (models.UserModel, bool)
	CheckIfUserNamePasswordExist(data UserData, database *Database) (models.UserModel, bool)
}

// Ошибка - не в сессии
var notInSession = errors.New("not in session")

// Структура сессия
type UserSession struct {
	Store *sessions.CookieStore
}

// Создание сессии
func CreateNewUserSession() UserSession {
	return UserSession{
		Store: sessions.NewCookieStore(securecookie.GenerateRandomKey(32)),
	}
}

// Добавление пользователя в сессию
func (userSession *UserSession) LoginUser(w http.ResponseWriter, r *http.Request, dataBase *Database, data *UserData) {
	session, _ := userSession.Store.Get(r, "cookie-name")

	var user, isExist = userSession.CheckIfUserNameExist(*data, dataBase)
	if isExist {
		log.Println("User: ", user.Name, " exist")
		session.Values["authenticated"] = true
		session.Values["id"] = user.ID
		session.Values["name"] = user.Name
		log.Println(session.Values["id"].(int))
	} else {
		log.Println("User not found")
	}
	session.Save(r, w)
}

// Удаление пользователя из сессии
func (userSession *UserSession) LogoutUser(w http.ResponseWriter, r *http.Request) {
	session, _ := userSession.Store.Get(r, "cookie-name")

	session.Values["authenticated"] = false
	session.Values["id"] = nil
	session.Values["name"] = nil

	session.Save(r, w)
}

// Проверка находится ли пользователь в сессии
func (userSession *UserSession) IsUserLogin(r *http.Request) bool {
	session, _ := userSession.Store.Get(r, "cookie-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return false
	}
	log.Println("User authenticated")
	return true
}

// Получение идентификатора и имени текущего пользователя в сессии
func (userSession *UserSession) GetCurrentUserIdName(r *http.Request) (int, string, error) {
	if userSession.IsUserLogin(r) {
		session, _ := userSession.Store.Get(r, "cookie-name")
		log.Println("User in session")
		return session.Values["id"].(int), session.Values["name"].(string), nil
	}
	return 0, "", notInSession
}

// Проверка существует ли пользователь с именем
func (userSession *UserSession) CheckIfUserNameExist(data UserData, database *Database) (models.UserModel, bool) {
	var user, err = database.GetUserByName(data.Name)
	if err == nil {
		log.Println("User exist in database")
		return user, true
	} else {
		log.Println("User not exist in database")
		return user, false
	}
}

// Проверка существует ли пользователь с таким именем и паролем
func (userSession *UserSession) CheckIfUserNamePasswordExist(data UserData, database *Database) (models.UserModel, bool) {
	var user, err = database.GetUserByNameAndPassword(data.Name, data.Password)
	if err == nil {
		return user, true
	} else {
		return user, false
	}
}
