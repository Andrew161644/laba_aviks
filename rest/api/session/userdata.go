package session

import "github.com/Andrew161644/avicks_laba/api/database/models"

type UserData struct {
	Name     string
	Password string
}

func GetUserData(user models.UserModel) UserData {
	return UserData{
		Name: user.Name,
	}
}
