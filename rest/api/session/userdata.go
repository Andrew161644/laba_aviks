package session

import "github.com/Andrew161644/avicks_laba/api/database/models"

type UserData struct {
	Name     string
	Password string
}

func ToUserData(user models.UserModel) UserData {
	return UserData{
		Name: user.Name,
	}
}
func IsValid(userdata UserData) bool {
	return len(userdata.Password) >= 5 && len(userdata.Name) > 1 && len(userdata.Name) <= 10 && len(userdata.Password) <= 8
}
